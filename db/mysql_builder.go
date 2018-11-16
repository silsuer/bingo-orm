package db

import (
	"database/sql"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

const (
	ModelTag = "json"
)

// 空接口
type emptyInterface struct {
	typ  *struct{}
	word unsafe.Pointer
}

type MysqlBuilder struct {
	connector       IConnector
	Builder                             // 塞入一个构建器
	bindings        map[string][]string // 绑定的操作符与列名之间的映射
	columns         []string            // 列名()
	distinct        bool                // 是否用到了去重查询
	distinctColumns []string            // 唯一的列
	from            string              // 表名
	joins           string              // 连接
	wheres          []*whereCondition   // where的数组
	groups          []string            // 组
	havings         []string            // group by 之后的操作
	orders          []string            // 排序
	limit           int                 // 限制
	offset          int                 // 偏移
	unions          []IBuilder          // 联合
	unionLimit      string
	unionOffset     string
	unionOrders     string
	lock            bool
	operator        int // 操作符
}

const (
	SelectOperator = iota
	UpdateOperator
	DeleteOperator
	InsertOperator
)

func NewMysqlBuilder() IBuilder {
	b := &MysqlBuilder{}
	b.operator = SelectOperator
	return b
}

// 去重
func (m *MysqlBuilder) Distinct(columns ...string) IBuilder {
	m.distinct = true
	m.distinctColumns = columns
	return m
}

// 删除某些数据
func (m *MysqlBuilder) Delete() IResult {
	// delete from table where ...
	whereSql, values := whereAssembly(m.wheres)
	s := `delete from ` + m.from + whereSql
	stmt, err := m.GetConn().Prepare(s)
	res := new(MysqlResult)
	if err != nil {
		res.errors = append(res.errors, err)
		return res
	}
	r, err := stmt.Exec(values...)
	if err != nil {
		res.errors = append(res.errors, err)
	}
	res.result = r
	return res
}

// 单独更新某个域
func (m *MysqlBuilder) UpdateField(column string, value interface{}) IResult {
	var values []interface{}
	values = append(values, value)
	whereSql, vals := whereAssembly(m.wheres)
	values = append(values, vals...)
	s := `update ` + m.from + ` set ` + column + `=?` + whereSql
	stmt, err := m.GetConn().Prepare(s)
	res := new(MysqlResult)
	if err != nil {
		res.errors = append(res.errors, err)
		return res
	}
	r, err := stmt.Exec(values...)
	if err != nil {
		res.errors = append(res.errors, err)
	}
	res.result = r
	return res
}

func (m *MysqlBuilder) UpdateModel(model interface{}) IResult {
	// update xxx set a=x,b=x where ...
	var fields []string
	var values []interface{}
	t := reflect.TypeOf(model).Elem()
	structPtr := (*emptyInterface)(unsafe.Pointer(&model)).word
	originPtr := uintptr(structPtr)
	var v interface{}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldPtr := originPtr + field.Offset // 根据偏移量获取当前属性的指针
		switch field.Type.Name() {
		case "int":
			v = *((*int)(unsafe.Pointer(fieldPtr)))
		case "string":
			v = *((*string)(unsafe.Pointer(fieldPtr)))
		}
		tag := field.Tag.Get(ModelTag)
		if tag != "" {
			fields = append(fields, tag+`=?`)
			values = append(values, v)
		}
	}

	whereSql, vals := whereAssembly(m.wheres)
	values = append(values, vals...)

	s := `update ` + m.from + ` set ` + strings.Join(fields, ",") + whereSql
	stmt, err := m.GetConn().Prepare(s)
	res := new(MysqlResult)
	if err != nil {
		res.errors = append(res.errors, err)
		return res
	}

	r, err := stmt.Exec(values...)
	if err != nil {
		res.errors = append(res.errors, err)
	}
	res.result = r
	return res
}

func (m *MysqlBuilder) UpdateMap(mm map[string]interface{}) IResult {
	// update table xxx set a=x,b=3 where a = 1
	var fields []string
	var values []interface{}
	for k := range mm {
		fields = append(fields, k+`=?`)
		values = append(values, mm[k])
	}
	whereSql, vals := whereAssembly(m.wheres)
	s := `update ` + m.from + ` set ` + strings.Join(fields, ",") + whereSql
	values = append(values, vals...)
	res := new(MysqlResult)
	stmt, err := m.GetConn().Prepare(s)
	if err != nil {
		res.errors = append(res.errors, err)
		return res
	}
	r, err := stmt.Exec(values...)
	if err != nil {
		res.errors = append(res.errors, err)
	}

	res.result = r
	return res
}

func whereAssembly(w []*whereCondition) (string, []interface{}) {
	// where a is null
	s := ` where `
	var cols []string
	var vals []interface{}
	for k := range w {
		// a > 0
		if w[k].columnFunc != "" {
			cols = append(cols, w[k].columnFunc+`(`+w[k].columnName+`)`+w[k].operator+`?`)
		} else {
			cols = append(cols, w[k].columnName+w[k].operator+`?`)
		}
		vals = append(vals, w[k].value)
	}
	return s + strings.Join(cols, ` and `), vals // 返回where语句和占位符
}

func (m *MysqlBuilder) GetConn() *sql.DB {
	return m.connector.GetConn()
}

func (m *MysqlBuilder) SetConn(connector IConnector) IBuilder {
	m.connector = connector
	return m
}

type MysqlResult struct {
	Result // 结果 重写这个接口下的所有方法
	errors []error
	result sql.Result
	row    *sql.Row
	rows   *sql.Rows
}

func (mr *MysqlResult) ToStringMapList() []map[string]string {
	// 将结果集生成map，并返回
	column, _ := mr.rows.Columns()
	values := make([][]byte, len(column)) // 每个列的值
	scans := make([]interface{}, len(column))
	for i := range values {
		scans[i] = &values[i]
	}

	var results []map[string]string
	for mr.rows.Next() { // 循环，让游标往下移动
		if err := mr.rows.Scan(scans...); err != nil {
			panic(err)
		}
		row := make(map[string]string) // 每行数据
		for k := range values {
			key := column[k]
			row[key] = string(values[k])
		}
		results = append(results, row)
	}
	return results
}

func (mr *MysqlResult) ToMapList() []map[string]interface{} {
	// 将结果集生成map，并返回
	column, _ := mr.rows.Columns()
	values := make([]interface{}, len(column)) // 每个列的值
	scans := make([]interface{}, len(column))
	for i := range values {
		scans[i] = &values[i]
	}

	var results []map[string]interface{}
	for mr.rows.Next() { // 循环，让游标往下移动
		if err := mr.rows.Scan(scans...); err != nil {
			panic(err)
		}
		row := make(map[string]interface{}) // 每行数据
		for k := range values {
			key := column[k]
			row[key] = values[k]
		}
		results = append(results, row)
	}
	return results
}

func (mr *MysqlResult) GetResult() sql.Result {
	return mr.result
}

func (mr *MysqlResult) GetRows() *sql.Rows {
	return mr.rows
}

func (mr *MysqlResult) GetRow() *sql.Row {
	return mr.row
}

func (mr *MysqlResult) GetErrors() []error {
	return mr.errors
}

func (m *MysqlBuilder) Table(tableName string) IBuilder {
	m.from = tableName
	return m
}

// 传入map切片，批量插入数据
func (m *MysqlBuilder) InsertManyMap(dm []map[string]interface{}) IResult {
	if len(dm) == 0 {
		return &MysqlResult{}
	}
	var columns []string // 用来记录要prepare的数据列
	var values []interface{}
	var valuesSlice []string
	// 列的顺序
	for k := range dm[0] {
		columns = append(columns, k)
	}
	// 值顺序
	for index := range dm {
		var prepares []string
		for key := range dm[index] {
			values = append(values, dm[index][key]) // 要prepare得数据
			prepares = append(prepares, "?")        //  prepare 占位符
		}
		valuesSlice = append(valuesSlice, `(`+strings.Join(prepares, ",")+`)`)
	}
	// 拼接字符串
	s := `insert into ` + m.from + ` (` + strings.Join(columns, ",") + `) values ` + strings.Join(valuesSlice, ",")

	res := new(MysqlResult)
	// 执行代码
	stmt, err := m.GetConn().Prepare(s)
	if err != nil {
		res.errors = append(res.errors, err)
		return res
	}
	r, err := stmt.Exec(values...)
	res.result = r
	if err != nil {
		res.errors = append(res.errors, err)
		return res
	}
	return res
}

// 传入多个结构体，批量插入数据
func (m *MysqlBuilder) InsertManyModels(models []interface{}) IResult {
	if len(models) == 0 {
		return &MysqlResult{}
	}

	var columns []string // 用来记录要prepare的数据列
	var fieldNames []string
	var values []interface{}
	var valuesSlice []string

	// 分析第一个模型，获取列名
	t := reflect.TypeOf(models[0]).Elem()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		columns = append(columns, field.Tag.Get(ModelTag))
		fieldNames = append(fieldNames, field.Name)
	}

	// 整体遍历
	for k := range models {
		var prepares []string
		tt := reflect.TypeOf(models[k]).Elem()
		structPtr := (*emptyInterface)(unsafe.Pointer(&models[k])).word
		originPtr := uintptr(structPtr)

		for kk := range fieldNames {
			field, _ := tt.FieldByName(fieldNames[kk])
			fieldPtr := originPtr + field.Offset
			var v interface{}
			switch field.Type.Name() {
			case "int":
				v = *((*int)(unsafe.Pointer(fieldPtr)))
			case "string":
				v = *((*string)(unsafe.Pointer(fieldPtr)))
			}
			// 根据指针偏移量获取值
			values = append(values, v)
			prepares = append(prepares, "?")
		}
		valuesSlice = append(valuesSlice, `(`+strings.Join(prepares, ",")+`)`)
	}

	// 拼接字符串
	s := `insert into ` + m.from + ` (` + strings.Join(columns, ",") + `) values ` + strings.Join(valuesSlice, ",")

	res := new(MysqlResult)
	// 执行代码
	stmt, err := m.GetConn().Prepare(s)
	if err != nil {
		res.errors = append(res.errors, err)
		return res
	}
	r, err := stmt.Exec(values...)
	res.result = r
	if err != nil {
		res.errors = append(res.errors, err)
		return res
	}
	return res
}

// 使用结构体插入数据
func (m *MysqlBuilder) InsertModel(model interface{}) IResult {

	var columns []string // 用来记录要prepare的数据列
	var values []interface{}
	var prepares []string

	// 获取所有的json标记，根据标记获取值
	t := reflect.TypeOf(model).Elem()
	structPtr := (*emptyInterface)(unsafe.Pointer(&model)).word
	originPtr := uintptr(structPtr)
	var v interface{}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldPtr := originPtr + field.Offset // 根据偏移量获取当前属性的指针
		switch field.Type.Name() {
		case "int":
			v = *((*int)(unsafe.Pointer(fieldPtr)))
		case "string":
			v = *((*string)(unsafe.Pointer(fieldPtr)))
		}

		tag := field.Tag.Get(ModelTag)
		if tag != "" {
			columns = append(columns, tag)
			prepares = append(prepares, "?")
			values = append(values, v)
		}
	}

	s := `insert into ` + m.from + ` (` + strings.Join(columns, ",") + `) values(` + strings.Join(prepares, ",") + `)`
	res := new(MysqlResult)
	// 执行插入
	r, err := m.GetConn().Exec(s, values...)
	if err != nil {
		res.errors = append(res.errors, err)
		return res
	}
	res.result = r
	return res
}

// 使用map 插入数据
func (m *MysqlBuilder) InsertMap(d map[string]interface{}) IResult {
	// insert into table_name (c1,c2,c3) values (v1,v2,v3)
	var columns []string // 用来记录要prepare的数据列
	var values []interface{}
	var prepares []string
	//sql = `insert into ` + m.from + ` ()`
	for k := range d {
		columns = append(columns, k)
		values = append(values, d[k])
		prepares = append(prepares, "?")
	}
	s := `insert into ` + m.from + ` (` + strings.Join(columns, ",") + `) values(` + strings.Join(prepares, ",") + `)`

	res := new(MysqlResult)
	// 执行插入
	//stmt, err := m.GetConn().Prepare(s)
	r, err := m.GetConn().Exec(s, values...)
	if err != nil {
		res.errors = append(res.errors, err)
		return res
	}
	res.result = r
	return res
}

func (m *MysqlBuilder) OrderBy(column string, sort string) IBuilder {
	m.orders = append(m.orders, column+` `+sort) // 添加排序规则
	return m
}

func (m *MysqlBuilder) OrderByDesc(column string) IBuilder {
	m.orders = append(m.orders, column+` desc`)
	return m
}

// 获取结果集
func (m *MysqlBuilder) Get() IResult {
	var values []interface{}
	// 拼接当前的数据
	s, vals := selectAssembly(m)
	values = append(values, vals...)

	// 执行
	rows, err := m.GetConn().Query(s, values...)
	res := new(MysqlResult)
	if err != nil {
		res.errors = append(res.errors, err)
	}
	res.rows = rows
	return res
}

func selectAssembly(m *MysqlBuilder) (string, []interface{}) {
	// select * from table where xxxx order by xxx  group by xxx having xxx limit x offset xxx join xxx union ...

	/**

	SELECT
	   [ALL | DISTINCT | DISTINCTROW ]
		 [HIGH_PRIORITY]
		 [STRAIGHT_JOIN]
		 [SQL_SMALL_RESULT] [SQL_BIG_RESULT] [SQL_BUFFER_RESULT]
		 [SQL_CACHE | SQL_NO_CACHE] [SQL_CALC_FOUND_ROWS]
	   select_expr [, select_expr ...]
	   [FROM table_references
		 [PARTITION partition_list]
	   [WHERE where_condition]
	   [GROUP BY {col_name | expr | position}
		 [ASC | DESC], ... [WITH ROLLUP]]
	   [HAVING where_condition]
	   [ORDER BY {col_name | expr | position}
		 [ASC | DESC], ...]
	   [LIMIT {[offset,] row_count | row_count OFFSET offset}]
	   [PROCEDURE procedure_name(argument_list)]
	   [INTO OUTFILE 'file_name'
		   [CHARACTER SET charset_name]
		   export_options
		 | INTO DUMPFILE 'file_name'
		 | INTO var_name [, var_name]]
	   [FOR UPDATE | LOCK IN SHARE MODE]]

	*/

	var values []interface{}

	s := `select `

	// distinct
	if m.distinct { // 存在distinct 列
		m.columns = m.distinctColumns // 如果是distinct，将会忽略前面出现的所有数据
		s += ` distinct `
	}
	// 拼接列
	if len(m.columns) == 0 {
		s += ` * `
	} else {
		s += strings.Join(m.columns, ",")
	}

	// 拼接表名
	s += ` from ` + m.from
	// 拼接where语句
	whereSql, vals := whereAssembly(m.wheres)
	s += whereSql
	values = append(values, vals...)

	// group by  having
	if len(m.groups) > 0 {
		s += ` group by ` + strings.Join(m.groups, ",")
	}
	if len(m.havings) > 0 {
		s += ` having ` + strings.Join(m.havings, ",")
	}

	// 拼接order by
	if len(m.orders) > 0 {
		s += ` order by ` + strings.Join(m.orders, ",")
	}

	// 拼接limit
	if m.limit != 0 {
		s += ` limit ` + strconv.Itoa(m.limit)
	}
	// 拼接offset
	if m.offset != 0 {
		s += ` offset ` + strconv.Itoa(m.offset)
	}

	return s, values
}

//func selectAssembly(m *MysqlBuilder) string {
//	// 根据传入的builder，构建数据
//	// select columns from table
//}

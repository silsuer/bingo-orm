package db

import (
	"strconv"
	"strings"
)

const (
	TypeInteger              = "int"
	TypeBigInteger           = "bigint"
	TypeString               = "varchar"
	TypeChar                 = "char"
	TypeBlob                 = "blob"
	TypeBool                 = "boolean"
	TypeDate                 = "date"
	TypeDateTime             = "datetime"
	TypeTimestamp            = "timestamp"
	TypeTime                 = "time"
	TypeDecimal              = "decimal"
	TypeDouble               = "double"
	TypeFloat                = "float"
	TypeGeometry             = "geometry"
	TypeGeometryCollection   = "geometrycollection"
	TypeJson                 = "json"
	TypeLineString           = "linestring"
	TypeLongText             = "longtext"
	TypeMediumInteger        = "mediumint"
	TypeTinyInteger          = "tinyint"
	TypeSmallInteger         = "smallint"
	TypeEnum                 = "enum"
	TypeMediumText           = "mediumtext"
	TypeText                 = "text"
	TypeMultiLineString      = "multilinestring"
	TypeMultiPoint           = "multipoint"
	TypeMultiPolygon         = "multipolygon"
	TypePoint                = "point"
	TypePolygon              = "polygon"
	TypeYear                 = "year"
	TypeStringDefaultLength  = 255
	TypeIntegerDefaultLength = 11

	CreateDefaultType = iota
	CreateIfNotExists
	AlterTable
	CreateTable
	DropTable
	ChangeColumn
	DropColumn
	AddColumn
	RenameColumn
)

// mysql实现 blueprint
type MysqlBlueprint struct {
	Blueprint
	name              string // 表名
	engine            string
	columns           []*MysqlColumn
	currentCol        *MysqlColumn
	charset           string
	indexList         []string   // 普通索引
	combineIndexList  [][]string // 组合索引数组
	uniqueIndexList   []string   // 唯一索引
	fulltextIndexList []string   // 全文索引
	primaryKey        string     // 主键索引
	operator          int        // 操作类型，是创建表还是更改表结构
}

type MysqlColumn struct {
	name             string // 列名
	columnType       string // 列类型
	defaultValue     string // 默认值
	defaultType      bool   // 默认值的类型 , true 是int，false是字符串
	comment          string // 备注
	nullable         bool   // 是否允许为null
	length           int    // 列的长度
	unsigned         bool   // 是否有符号
	autoIncrement    bool   // 是否自增
	originColumnName string // 原列名（只在重命名列时用到）
	operator         int    // 列操作符
}

func NewColumn(mb *MysqlBlueprint, name string) *MysqlColumn {
	c := new(MysqlColumn)
	c.nullable = false
	c.name = name
	c.operator = AddColumn
	c.unsigned = false // 默认都是有符号的
	c.autoIncrement = false
	c.defaultType = false // 默认值是 字符串
	mb.columns = append(mb.columns, c)
	mb.currentCol = c
	return c
}

func checkLength(length []int) {
	if len(length) > 1 {
		panic("too many arguments in String() function")
	}
}

func (mb *MysqlBlueprint) Binary(columnName string) IBlueprint {
	// 建立一个 blob类型的字段
	column := NewColumn(mb, columnName)
	column.columnType = TypeBlob
	return mb
}

func (mb *MysqlBlueprint) Increments(columnName string, length ...int) IBlueprint {

	checkLength(length)

	// 创建一个列
	column := NewColumn(mb, columnName)
	column.length = TypeIntegerDefaultLength

	if len(length) == 1 {
		column.length = length[0]
	}

	column.columnType = TypeInteger
	column.autoIncrement = true // 自动递增
	column.unsigned = true      // 无符号
	column.defaultType = true   // 默认值改成 数字
	mb.PrimaryKey(columnName)   // 设置主键
	//mb.columns = append(mb.columns, column)
	//mb.currentCol = column // 设置当前的列名
	return mb
}

// 使用 int 10 存储 ip地址，使用 inet_ntoa() 将数字转成ip地址，使用 inet_aton() 将ip地址字符串转成数字
func (mb *MysqlBlueprint) IpAddress(columnName string) IBlueprint {
	return mb.Integer(columnName, 10)
}

func (mb *MysqlBlueprint) LineString(columnName string) IBlueprint {
	c := NewColumn(mb, columnName)
	c.columnType = TypeLineString
	return mb
}

func (mb *MysqlBlueprint) LongText(columnName string) IBlueprint {
	c := NewColumn(mb, columnName)
	c.columnType = TypeLongText
	return mb
}

//func (mb *MysqlBlueprint) Float(columnName string, lengthPrecision ...int) IBlueprint {
//
//}

func (mb *MysqlBlueprint) Enum(columnName string, enum ...string) IBlueprint {
	// name enum('a','b')
	c := NewColumn(mb, columnName)
	c.columnType = TypeEnum

	var tmpSlice []string
	for k := range enum {
		tmpSlice = append(tmpSlice, `"`+enum[k]+`"`)
	}
	c.columnType += `(` + strings.Join(tmpSlice, ",") + `)`
	return mb
}

func (mb *MysqlBlueprint) MediumInteger(columnName string, length ...int) IBlueprint {
	checkLength(length)
	c := NewColumn(mb, columnName)
	c.columnType = TypeMediumInteger
	c.defaultValue = strconv.Itoa(TypeIntegerDefaultLength)
	c.defaultType = true
	if len(length) == 1 {
		c.length = length[0]
	}

	return mb
}

func (mb *MysqlBlueprint) MediumIncrements(columnName string, length ...int) IBlueprint {
	checkLength(length)

	c := NewColumn(mb, columnName)
	c.columnType = TypeMediumInteger
	c.defaultType = true
	c.unsigned = true
	c.autoIncrement = true
	c.length = TypeIntegerDefaultLength

	if len(length) == 1 {
		c.length = length[0]
	}

	mb.PrimaryKey(columnName)
	return mb
}

// 存储json字符串
func (mb *MysqlBlueprint) Json(columnName string) IBlueprint {
	c := NewColumn(mb, columnName)
	c.columnType = TypeJson
	return mb
}

func (mb *MysqlBlueprint) GeometryCollection(columnName string) IBlueprint {
	c := NewColumn(mb, columnName)
	c.columnType = TypeGeometryCollection
	return mb
}

func (mb *MysqlBlueprint) Geometry(columnName string) IBlueprint {
	column := NewColumn(mb, columnName)
	column.columnType = TypeGeometry
	return mb
}

func (mb *MysqlBlueprint) Float(columnName string, lengthPrecision ...int) IBlueprint {
	c := NewColumn(mb, columnName)
	c.columnType = TypeFloat
	switch len(lengthPrecision) {
	case 0:
	case 2:
		c.columnType += `(` + strconv.Itoa(lengthPrecision[0]) + `,` + strconv.Itoa(lengthPrecision[1]) + `)`
	default:
		panic("wrong arguments number in Float() function. need 2 arguments: length and precision.")
	}

	return mb
}

func (mb *MysqlBlueprint) MediumText(columnName string) IBlueprint {
	c := NewColumn(mb, columnName)
	c.columnType = TypeMediumText
	return mb
}

// multilinestring
func (mb *MysqlBlueprint) MultiLineString(columnName string) IBlueprint {
	c := NewColumn(mb, columnName)
	c.columnType = TypeMultiLineString
	return mb
}

func (mb *MysqlBlueprint) MultiPoint(columnName string) IBlueprint {
	c := NewColumn(mb, columnName)
	c.columnType = TypeMultiPoint
	return mb
}

func (mb *MysqlBlueprint) Timestamp(columnName string) IBlueprint {
	c := NewColumn(mb, columnName)
	c.columnType = TypeTimestamp
	return mb
}

func (mb *MysqlBlueprint) TinyIncrements(columnName string, length ...int) IBlueprint {
	mb.TinyInteger(columnName, length...)
	mb.PrimaryKey(columnName)
	mb.currentCol.unsigned = true
	mb.currentCol.autoIncrement = true
	return mb
}

func (mb *MysqlBlueprint) Polygon(columnName string) IBlueprint {
	c := NewColumn(mb, columnName)
	c.columnType = TypePolygon
	return mb
}

func (mb *MysqlBlueprint) SmallInteger(columnName string, length ...int) IBlueprint {

	checkLength(length)
	c := NewColumn(mb, columnName)
	c.columnType = TypeSmallInteger
	c.defaultType = true
	if len(length) == 1 {
		c.length = length[0]
	}
	return mb
}

func (mb *MysqlBlueprint) UnsignedInteger(columnName string, length ...int) IBlueprint {
	mb.Integer(columnName, length...)
	mb.currentCol.unsigned = true
	return mb
}

func (mb *MysqlBlueprint) UnsignedTinyInteger(columnName string, length ...int) IBlueprint {
	mb.TinyInteger(columnName, length...)
	mb.currentCol.unsigned = true
	return mb
}

func (mb *MysqlBlueprint) Year(columnName string) IBlueprint {
	c := NewColumn(mb, columnName)
	c.columnType = TypeYear
	return mb
}

//func (mb *MysqlBlueprint) Uuid(columnName string) IBlueprint {
//	mb.String(columnName, 36).Default("uuid()")
//	return mb
//}

func (mb *MysqlBlueprint) UnsignedSmallInteger(columnName string, length ...int) IBlueprint {
	mb.SmallInteger(columnName, length...)
	mb.currentCol.unsigned = true
	return mb
}

func (mb *MysqlBlueprint) UnsignedMediumInteger(columnName string, length ...int) IBlueprint {
	mb.MediumInteger(columnName, length...)
	mb.currentCol.unsigned = true
	return mb
}

func (mb *MysqlBlueprint) UnsignedDecimal(columnName string, length int, precision int) IBlueprint {
	mb.Decimal(columnName, length, precision)
	mb.currentCol.unsigned = true
	return mb
}

func (mb *MysqlBlueprint) Text(columnName string) IBlueprint {
	c := NewColumn(mb, columnName)
	c.columnType = TypeText
	return mb
}

// 删除当前列
func (mb *MysqlBlueprint) Drop() IBlueprint {
	mb.currentCol.operator = DropColumn
	return mb
}

func (mb *MysqlBlueprint) Change() IBlueprint {
	mb.currentCol.operator = ChangeColumn
	return mb
}

// 重命名列
func (mb *MysqlBlueprint) RenameColumn(originColumn string) IBlueprint {
	mb.currentCol.operator = RenameColumn
	mb.currentCol.originColumnName = originColumn // 原列名
	return mb
}

func (mb *MysqlBlueprint) Time(columnName string) IBlueprint {
	c := NewColumn(mb, columnName)
	c.columnType = TypeTime
	return mb
}

func (mb *MysqlBlueprint) UnsignedBigInteger(columnName string, length ...int) IBlueprint {
	mb.BigInteger(columnName, length...)
	mb.currentCol.unsigned = true
	return mb
}

func (mb *MysqlBlueprint) SoftDeletes() IBlueprint {
	mb.Timestamp("deleted_at").Nullable()
	return mb
}

func (mb *MysqlBlueprint) SmallIncrements(columnName string, length ...int) IBlueprint {
	mb.SmallInteger(columnName, length...)
	mb.currentCol.autoIncrement = true
	mb.currentCol.unsigned = true
	return mb
}

func (mb *MysqlBlueprint) RememberToken() IBlueprint {
	mb.String("remember_token", 100).Nullable()
	return mb
}

func (mb *MysqlBlueprint) Point(columnName string) IBlueprint {
	c := NewColumn(mb, columnName)
	c.columnType = TypePoint
	return mb
}

func (mb *MysqlBlueprint) TinyInteger(columnName string, length ...int) IBlueprint {

	checkLength(length)

	c := NewColumn(mb, columnName)
	c.columnType = TypeTinyInteger
	c.defaultType = true

	if len(length) == 1 {
		c.length = length[0]
	}

	return mb
}

// 自动创建 created_at updated_at
func (mb *MysqlBlueprint) Timestamps() IBlueprint {
	mb.NullableTimestamp("created_at")
	mb.NullableTimestamp("updated_at")
	return mb
}

func (mb *MysqlBlueprint) NullableTimestamp(columnName string) IBlueprint {
	mb.Timestamp(columnName).Nullable()
	return mb
}

func (mb *MysqlBlueprint) MultiPolygon(columnName string) IBlueprint {
	c := NewColumn(mb, columnName)
	c.columnType = TypeMultiPolygon
	return mb
}

func (mb *MysqlBlueprint) Double(columnName string, lengthPrecision ...int) IBlueprint {
	column := NewColumn(mb, columnName)
	column.columnType = TypeDouble

	switch len(lengthPrecision) {
	case 0:
	case 2:
		column.columnType += `(` + strconv.Itoa(lengthPrecision[0]) + `,` + strconv.Itoa(lengthPrecision[1]) + `)`
	default:
		panic("wrong arguments number in Double() function. need 2 arguments: length and precision.")
	}

	return mb
}

func (mb *MysqlBlueprint) Decimal(columnName string, length int, precision int) IBlueprint {
	column := NewColumn(mb, columnName)
	column.columnType = TypeDecimal + `(` + strconv.Itoa(length) + `,` + strconv.Itoa(precision) + `)`
	return mb
}

func (mb *MysqlBlueprint) Nullable() IBlueprint {
	mb.currentCol.nullable = true
	return mb
}

func (mb *MysqlBlueprint) Comment(comment string) IBlueprint {
	mb.currentCol.comment = comment
	return mb
}

func (mb *MysqlBlueprint) Default(def interface{}) IBlueprint {
	if mb.currentCol.defaultType { // 需要 int 型默认值
		if v, ok := def.(int); ok {
			mb.currentCol.defaultValue = strconv.Itoa(v)
		} else {
			panic("wrong default value type (need int) : column " + mb.currentCol.name)
		}
	} else {
		// 需要 字符串型 默认值
		if v, ok := def.(string); ok {
			mb.currentCol.defaultValue = v
		} else {
			panic("wrong default value type (need string) : column " + mb.currentCol.name)
		}
	}

	return mb
}

func (mb *MysqlBlueprint) String(columnName string, length ...int) IBlueprint {

	checkLength(length)

	column := NewColumn(mb, columnName)
	column.length = TypeStringDefaultLength

	if len(length) == 1 {
		column.length = length[0]
	}

	// 字符串类型
	column.columnType = TypeString
	//mb.columns = append(mb.columns, column)
	//mb.currentCol = column
	return mb
}

func (mb *MysqlBlueprint) Boolean(columnName string) IBlueprint {
	// 建立一个bool类型的字段
	column := NewColumn(mb, columnName)
	column.columnType = TypeBool
	column.defaultType = true // 默认值是int型
	return mb
}

func (mb *MysqlBlueprint) DateTime(columnName string) IBlueprint {
	column := NewColumn(mb, columnName)
	column.columnType = TypeDateTime
	return mb
}

func (mb *MysqlBlueprint) Date(columnName string) IBlueprint {
	column := NewColumn(mb, columnName)
	column.columnType = TypeDate
	return mb
}

func (mb *MysqlBlueprint) Char(columnName string, length int) IBlueprint {
	column := NewColumn(mb, columnName)
	column.columnType = TypeChar
	return nil
}

func (mb *MysqlBlueprint) Integer(columnName string, length ...int) IBlueprint {
	if len(length) > 1 {
		panic("too many arguments in Integer() function.")
	}

	column := NewColumn(mb, columnName)
	column.defaultType = true
	column.length = TypeIntegerDefaultLength

	if len(length) == 1 {
		column.length = length[0]
	}

	column.columnType = TypeInteger
	//mb.columns = append(mb.columns, column)
	//mb.currentCol = column
	return mb
}

func (mb *MysqlBlueprint) Charset(charset string) IBlueprint {
	mb.charset = charset
	return mb
}

func (mb *MysqlBlueprint) IntegerWithLength(columnName string, length int) IBlueprint {
	column := NewColumn(mb, columnName)
	column.length = length
	column.columnType = TypeInteger
	//mb.columns = append(mb.columns, column)
	//mb.currentCol = column
	return mb
}

// 设置引擎
func (mb *MysqlBlueprint) Engine(engine string) IBlueprint {
	mb.engine = engine
	return mb
}

// 普通索引
func (mb *MysqlBlueprint) Index(columns ...string) {
	if len(columns) == 1 {
		mb.indexList = append(mb.indexList, columns[0])
	} else {
		mb.combineIndexList = append(mb.combineIndexList, columns)
	}
}

// 唯一索引
func (mb *MysqlBlueprint) UniqueIndex(column string) {
	mb.uniqueIndexList = append(mb.uniqueIndexList, column)
}

// 全文索引
func (mb *MysqlBlueprint) FullTextIndex(column string) {
	mb.fulltextIndexList = append(mb.fulltextIndexList, column)
}

func (mb *MysqlBlueprint) BigInteger(columnName string, length ...int) IBlueprint {
	checkLength(length)

	column := NewColumn(mb, columnName)
	column.columnType = TypeBigInteger
	column.defaultType = true
	if len(length) == 1 {
		column.defaultValue = strconv.Itoa(length[0])
	}
	//mb.columns = append(mb.columns, column)
	//mb.currentCol = column
	return mb
}

// 自增id 相当于 unsigned big integer
func (mb *MysqlBlueprint) BigIncreaments(columnName string, length ...int) IBlueprint {

	checkLength(length)

	column := NewColumn(mb, columnName)
	column.columnType = TypeBigInteger
	column.defaultType = true
	column.autoIncrement = true
	column.unsigned = true
	if len(length) == 1 {
		column.length = length[0]
	}

	//mb.columns = append(mb.columns, column)
	//mb.currentCol = column
	mb.PrimaryKey(columnName) // 设为主键
	return mb
}

func (mb *MysqlBlueprint) PrimaryKey(column string) {
	mb.primaryKey = column
}

// 修改表结构的方法
func alterAssembly(schema *MysqlBlueprint) []string {
	var sql []string
	for k := range schema.columns {
		switch schema.columns[k].operator {
		case AddColumn:
			sql = append(sql, `alter table `+schema.name+` add `+columnAssembly(schema.columns[k]))
		case ChangeColumn:
			sql = append(sql, `alter table `+schema.name+` modify `+columnAssembly(schema.columns[k]))
		case DropColumn:
			sql = append(sql, `alter table `+schema.name+` drop `+schema.columns[k].name)
		case RenameColumn:
			sql = append(sql, `alter table `+schema.name+` change `+schema.columns[k].originColumnName+` `+columnAssembly(schema.columns[k]))
		}
	}

	return sql
}

// 将表结构拼装成sql语句
func Assembly(createType int, schema *MysqlBlueprint) string {
	var sql string
	if createType == CreateDefaultType {
		sql = `CREATE TABLE ` + schema.name + "("
	} else if createType == CreateIfNotExists {
		sql = `CREATE TABLE IF NOT EXISTS ` + schema.name + "("
	}

	var columnSql []string
	// 拼接列
	for k := range schema.columns {
		columnSql = append(columnSql, columnAssembly(schema.columns[k]))
	}

	// 拼接索引
	// 拼接主键索引
	if schema.primaryKey != "" {
		columnSql = append(columnSql, primaryKeyAssembly(schema.primaryKey))
	}

	// 拼接普通索引
	for k := range schema.indexList {
		columnSql = append(columnSql, indexAssembly(schema.indexList[k]))
	}
	// 拼接唯一索引
	for k := range schema.uniqueIndexList {
		columnSql = append(columnSql, uniqueIndexAssembly(schema.uniqueIndexList[k]))
	}
	// 拼接组合索引
	for k := range schema.combineIndexList {
		columnSql = append(columnSql, combineIndexAssembly(schema.combineIndexList[k]))
	}
	// 拼接全文索引
	for k := range schema.fulltextIndexList {
		columnSql = append(columnSql, fulltextIndexAssembly(schema.fulltextIndexList[k]))
	}

	sql += strings.Join(columnSql, ",") + ")ENGINE=" + schema.engine

	if schema.charset != "" {
		sql += " DEFAULT CHARSET=" + schema.charset
	}

	return sql
}

// 主键索引
func primaryKeyAssembly(name string) string {
	//primary key (user_id)
	return `primary key (` + name + `)`
}

// 拼接普通索引
func indexAssembly(index string) string {
	return `index(` + index + `)`
}

// 拼接唯一索引
func uniqueIndexAssembly(index string) string {
	return `unique(` + index + `)`
}

// 拼接全文索引
func fulltextIndexAssembly(index string) string {
	return `fulltext(` + index + `)`
}

// 添加组合索引
func combineIndexAssembly(columns []string) string {
	return `index(` + strings.Join(columns, ",") + `)`
}

// 将列拼接成语句
func columnAssembly(column *MysqlColumn) string {
	var sql string

	// 列名 类型 （长度） 是否有符号 是否为空 是否自动递增 默认值 备注

	// 拼名称
	sql = column.name

	// 拼类型
	sql += ` ` + column.columnType

	// 拼类型长度
	if column.length != 0 {
		sql += ` (` + strconv.Itoa(column.length) + `) `
	}

	// 拼接是否有符号
	if column.unsigned { // 无符号
		sql += ` unsigned `
	}

	// 拼接是否允许为空
	if !column.nullable {
		sql += ` not null `
	}

	// 拼接自动递增
	if column.autoIncrement {
		sql += ` auto_increment `
	}

	// 拼接默认值
	if column.defaultValue != "" {
		// 拼接
		sql += ` default `
		if column.defaultType {
			sql += column.defaultValue
		} else {
			sql += `"` + column.defaultValue + `"`
		}
	}

	// 拼接备注
	if column.comment != "" {
		sql += ` comment "` + column.comment + `"`
	}
	return sql
}

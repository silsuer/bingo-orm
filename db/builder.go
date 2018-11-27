package db

import (
	"database/sql"
)

type IBuilder interface {
	SetConn(connector IConnector) IBuilder
	GetConn() *sql.DB
	Distinct(columns ...string) IBuilder     // 去重方法
	Where(condition ...interface{}) IBuilder // 条件语句
	//WhereCondition(columnName string, f func()) IBuilder // where语句，传入一个回调

	OrderByDesc(column string) IBuilder
	OrderBy(column string, sort string) IBuilder // 排序

	Get() IResult // 获取查询数据

	InnerJoin(tableName string) IBuilder // 内连接
	LeftJoin(tableName string) IBuilder  // 左连接
	RightJoin(tableName string) IBuilder // 右连接
	FullJoin(tableName string) IBuilder  // 完全连接

	On(conditions ...interface{}) IBuilder // 连接时候的条件

	//First() IResult
	//Find(id int) IResult
	//Pluck(column string) []interface{} // 传入一个列名，返回这个列的所有数据
	//Value(column string) interface{}   // 从单行中得到单独一个列的数据
	//Count(column string) IResult
	//Max(column string) IResult
	//Avg(column string) IResult
	//Min(column string) IResult
	//Sum(column string) IResult

	Table(tableName string) IBuilder                      // 设置表名
	InsertMap(d map[string]interface{}) IResult           // 传入一个map进行插入数据
	InsertModel(model interface{}) IResult                // 传入一个结构体进行插入数据,这个结构体必须要实现json标记
	InsertManyMap(dm []map[string]interface{}) IResult    // 传入一个map的切片
	InsertManyModels(models []interface{}) IResult        // 传入一个模型的数组
	UpdateMap(m map[string]interface{}) IResult           // 传入一个map，更新数据
	UpdateModel(model interface{}) IResult                // 传入一个结构体，更新数据
	UpdateField(column string, value interface{}) IResult // 单独更新某个域
	Delete() IResult                                      // 删除数据
	//UpdateManyMap(m []map[string]interface{}) IResult  // 传入一个map切片，批量更新数据
	//UpdateManyModels(models []interface{}) IResult     // 传入一个模型切片，批量更新数据
}

type IResult interface {
	ToModel() interface{}                // 转换为model
	GetRows() *sql.Rows                  // 获取多行数据
	GetRow() *sql.Row                    // 获取一行
	GetResult() sql.Result               // 获取结果
	GetErrors() []error                  // 获取查询过程中的错误
	ToMapList() []map[string]interface{} // 把查询结果转换为map
	//ToMap() map[string]interface{}
	ToStringMapList() []map[string]string // 把查询结果转换为字符串map
}

type Result struct {
}

func (r *Result) ToStringMapList() []map[string]string {
	return nil
}

func (r *Result) GetErrors() []error {
	return nil
}

func (r *Result) GetResult() sql.Result {
	return nil
}

func (r *Result) GetRow() *sql.Row {
	return nil
}

func (r *Result) GetRows() *sql.Rows {
	return nil
}

func (r *Result) ToModel() interface{} {
	return nil
}

func (r *Result) ToMapList() []map[string]interface{} {
	return nil
}

// 数据库/表 构建器
type ISchemaBuilder interface {
	SetConn(connector IConnector) ISchemaBuilder
	GetConn() *sql.DB

	// 创建数据库
	CreateDatabase(args ...string) (sql.Result, error)
	CreateDatabaseIfNotExists(args ...string) (sql.Result, error)
	// 创建数据表
	CreateTableIfNotExist(tableName string, call func(table IBlueprint)) error
	CreateTable(tableName string, call func(table IBlueprint)) error
	// 删除数据库
	DropDatabase(databaseName string) error
	// 删除数据表
	DropTable(tableName string) error
	// 清空数据库(删除数据库内所有表而不删除数据库)
	TruncateDatabase(databaseName string) error
	// 清空数据表
	TruncateTable(tableName string) error

	// 更改表结构
	Table(tableName string, call func(table IBlueprint)) error
	Transaction(t func(transaction ITransaction) error) error // 事务处理
}

type ITransaction interface {
	SetConn(connector IConnector)
	GetConn() IConnector
	Begin() error
	Rollback() error
	Commit() error
}

type Transaction struct {
}

func (t *Transaction) GetConn() IConnector {
	return nil
}

func (t *Transaction) SetConn(connector IConnector) {

}

func (t *Transaction) Begin() error {
	return nil
}

func (t *Transaction) Commit() error {
	return nil
}

func (t *Transaction) Rollback() error {
	return nil
}

type SchemaBuilder struct {
	connector IConnector
}

func (s *SchemaBuilder) Transaction(t func(transaction ITransaction) error) error {
	return nil
}

func (s *SchemaBuilder) Table(tableName string, call func(table IBlueprint)) error {
	return nil
}

func (s *SchemaBuilder) CreateTable(tableName string, call func(table IBlueprint)) error {
	return nil
}

func (s *SchemaBuilder) SetConn(connector IConnector) ISchemaBuilder {
	s.connector = connector
	return s
}

func (s *SchemaBuilder) GetConn() *sql.DB {
	return s.connector.GetConn()
}

func (s *SchemaBuilder) TruncateTable(tableName string) error {
	return nil
}

func (s *SchemaBuilder) TruncateDatabase(databaseName string) error {
	return nil
}

func (s *SchemaBuilder) DropTable(tableName string) error {
	return nil
}

func (s *SchemaBuilder) DropDatabase(databaseName string) error {
	return nil
}

func (s *SchemaBuilder) CreateTableIfNotExist(tableName string, call func(table IBlueprint)) error {
	return nil
}

func (s *SchemaBuilder) CreateDatabase(args ...string) (sql.Result, error) {
	return nil, nil
}

func (s *SchemaBuilder) CreateDatabaseIfNotExists(args ...string) (sql.Result, error) {
	return nil, nil
}

type Builder struct {
	//connector IConnector
}

func (b *Builder) On(condition ...interface{}) IBuilder {
	return nil
}

// 左连接
func (b *Builder) LeftJoin(tableName string) IBuilder {
	return nil
}

func (b *Builder) RightJoin(tableName string) IBuilder {
	return nil
}

func (b *Builder) FullJoin(tableName string) IBuilder {
	return nil
}

// 内连接 inner join
func (b *Builder) InnerJoin(tableName string) IBuilder {
	return nil
}

// order by
func (b *Builder) OrderBy(column string, sort string) IBuilder {
	return nil
}

func (b *Builder) OrderByDesc(column string) IBuilder {
	return nil
}

func (b *Builder) Delete() IResult {
	return nil
}

// 单独更新某个字段
func (b *Builder) UpdateField(column string, value interface{}) IResult {
	return nil
}

// 更新模型
func (b *Builder) UpdateModel(model interface{}) IResult {
	return nil
}

func (b *Builder) UpdateMap(m map[string]interface{}) IResult {
	return nil
}

func (b *Builder) Where(condition ...interface{}) IBuilder {
	return nil
}

// 传入模型来插入数据
func (b *Builder) InsertModel(model interface{}) IResult {
	return nil
}

func (b *Builder) InsertManyMap(dm []map[string]interface{}) IResult {
	return nil
}

func (b *Builder) InsertManyModels(models []interface{}) IResult {
	return nil
}

func (b *Builder) InsertMap(d map[string]interface{}) IResult {
	return nil
}

func (b *Builder) Distinct(columns ...string) IBuilder {
	return nil
}

func (b *Builder) Get() IResult {
	return nil
}

func (b *Builder) Table(tableName string) IBuilder {
	return nil
}

// 放入连接器
func (b *Builder) SetConn(connector IConnector) IBuilder {
	//b.connector = connector
	//return b
	return nil
}

func (b *Builder) GetConn() *sql.DB {
	//return b.connector.GetConn()
	return nil
}

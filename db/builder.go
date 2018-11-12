package db

import (
	"database/sql"
)

type IBuilder interface {
	SetConn(connector IConnector) IBuilder
	GetConn() *sql.DB
	Distinct(column string) IBuilder // 去重方法
	Get() (*sql.Rows, error)         //*sql.Rows  获取查询数据
	Table(tableName string) IBuilder // 设置表名
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
	connector IConnector
}

func (b *Builder) Distinct(column string) IBuilder {
	return nil
}

func (b *Builder) Get() (*sql.Rows, error) {
	return nil, nil
}

func (b *Builder) Table(tableName string) IBuilder {
	return nil
}

// 放入连接器
func (b *Builder) SetConn(connector IConnector) IBuilder {
	b.connector = connector
	return b
}

func (b *Builder) GetConn() *sql.DB {
	return b.connector.GetConn()
}

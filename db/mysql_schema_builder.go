package db

import (
	"database/sql"
	"errors"
	"strings"
)

type MysqlSchemaBuilder struct {
	SchemaBuilder
	connector IConnector
}

func NewMysqlSchemaBuilder() *MysqlSchemaBuilder {
	return &MysqlSchemaBuilder{}
}

func (ms *MysqlSchemaBuilder) SetConn(connector IConnector) ISchemaBuilder {
	ms.connector = connector
	return ms
}

func (ms *MysqlSchemaBuilder) GetConn() *sql.DB {
	return ms.connector.GetConn()
}

// mysql创建数据库
func (ms *MysqlSchemaBuilder) CreateDatabase(args ...string) (sql.Result, error) {
	if len(args) == 0 {
		return nil, errors.New(`the function CreateDatabase need 1 or 2 or 3 args. Use Case: CreateDatabase("db_name","gbk","gbk_chinese_ci")`)
	}
	dbName := args[0]              // 指定数据库名称
	dbCharset := "utf8"            // 指定字符集
	dbCollate := "utf8_general_ci" // 指定默认排序规则

	//var sql string
	switch len(args) {
	case 1:
	case 2:
		if strings.ToUpper(args[1]) == "UTF8" || strings.ToUpper(args[1]) == "UTF-8" {
			dbCharset = "utf8"
		} else if strings.ToUpper(args[1]) == "GBK" {
			dbCharset = "gbk"
			dbCollate = "gbk_chinese_ci"
		} else {
			// 除了这utf8和gbk这两种，其他的，必须指明排序规则
			return nil, errors.New("you must set your collate when create a new database")
		}
	case 3:
	default:
		return nil, errors.New(`too many arguments in the CreateDatabase function`)
	}

	// 执行操作
	stmt, err := ms.GetConn().Prepare(`CREATE DATABASE ` + dbName + ` DEFAULT CHARSET ` + dbCharset + ` COLLATE ` + dbCollate)
	if err != nil {
		return nil, err
	}
	return stmt.Exec()
}

func (ms *MysqlSchemaBuilder) CreateDatabaseIfNotExists(args ...string) (sql.Result, error) {
	if len(args) == 0 {
		return nil, errors.New(`the function CreateDatabase need 1 or 2 or 3 args. Use Case: CreateDatabase("db_name","gbk","gbk_chinese_ci")`)
	}
	dbName := args[0]              // 指定数据库名称
	dbCharset := "utf8"            // 指定字符集
	dbCollate := "utf8_general_ci" // 指定默认排序规则

	//var sql string
	switch len(args) {
	case 1:
	case 2:
		if strings.ToUpper(args[1]) == "UTF8" || strings.ToUpper(args[1]) == "UTF-8" {
			dbCharset = "utf8"
		} else if strings.ToUpper(args[1]) == "GBK" {
			dbCharset = "gbk"
			dbCollate = "gbk_chinese_ci"
		} else {
			// 除了这utf8和gbk这两种，其他的，必须指明排序规则
			return nil, errors.New("you must set your collate when create a new database")
		}
	case 3:
	default:
		return nil, errors.New(`too many arguments in the CreateDatabase function`)
	}

	// 执行操作
	stmt, err := ms.GetConn().Prepare(`CREATE DATABASE IF NOT EXISTS ` + dbName + ` DEFAULT CHARSET ` + dbCharset + ` COLLATE ` + dbCollate)
	if err != nil {
		return nil, err
	}
	return stmt.Exec()
}

const (
	SchemaDefaultEngine = "innodb"
)

func (ms *MysqlSchemaBuilder) CreateTableIfNotExist(tableName string, call func(table IBlueprint)) error {
	// 创建表
	schema := new(MysqlBlueprint)
	schema.engine = SchemaDefaultEngine // 默认引擎
	schema.name = tableName
	schema.operator = CreateTable
	call(schema)
	// 将schema拼接成sql语句
	// 调用完成，可以开始拼接数据了
	sql := Assembly(CreateIfNotExists, schema)
	stmt, err := ms.GetConn().Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	return err
}

func (ms *MysqlSchemaBuilder) Transaction(t func(transaction ITransaction) error) error {
	var err error
	transaction := new(MysqlTransaction)
	transaction.SetConn(ms.connector)
	err = transaction.Begin()
	if err != nil {
		return err
	}
	err = t(transaction)
	if err != nil {
		return err
	}
	err = transaction.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (ms *MysqlSchemaBuilder) CreateTable(tableName string, call func(table IBlueprint)) error {
	schema := new(MysqlBlueprint)
	schema.engine = SchemaDefaultEngine
	schema.name = tableName
	schema.operator = AlterTable
	call(schema)

	s := Assembly(CreateDefaultType, schema) // 拼装成语句
	//fmt.Println(sql)
	//return nil
	stmt, err := ms.GetConn().Prepare(s)
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	return err
}

// mysql事务处理
type MysqlTransaction struct {
	Transaction
	connector IConnector
}

func (mt *MysqlTransaction) SetConn(connector IConnector) {
	mt.connector = connector
}

func (mt *MysqlTransaction) GetConn() IConnector {
	return mt.connector
}

// 开启事务
func (mt *MysqlTransaction) Begin() error {
	_, err := mt.GetConn().GetConn().Exec(`START TRANSACTION`)
	return err
}

func (mt *MysqlTransaction) Commit() error {
	_, err := mt.GetConn().GetConn().Exec(`COMMIT`)
	return err
}

func (mt *MysqlTransaction) Rollback() error {
	stmt, err := mt.GetConn().GetConn().Prepare(`ROLLBACK`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec()
	return err
}

func (ms *MysqlSchemaBuilder) Table(tableName string, call func(table IBlueprint)) error {

	schema := new(MysqlBlueprint)
	schema.name = tableName
	schema.operator = AlterTable
	call(schema)
	// 拼接出sql数组 每一项都是一条sql语句,遍历并执行sql
	// 更改表结构的方法，返回拼接之后的数据
	s := alterAssembly(schema)

	var err error

	// 开启一个事务，事务执行失败将会全部回滚
	err = ms.Transaction(func(transaction ITransaction) error {
		for k := range s {
			// 遍历执行方法
			stmt, err := ms.GetConn().Prepare(s[k])

			if err != nil {
				return err
			}

			_, err = stmt.Exec()

			if err != nil {
				return err
			}
		}
		return nil
	})

	return err
}

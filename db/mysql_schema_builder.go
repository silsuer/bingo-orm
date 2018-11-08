package db

import (
	"errors"
	"strings"
	"database/sql"
	"fmt"
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
	call(schema)
	// 将schema拼接成sql语句
	// 调用完成，可以开始拼接数据了
	//var sql string
	fmt.Println(schema)
	return nil
}

func (ms *MysqlSchemaBuilder) CreateTable(tableName string, call func(table IBlueprint)) error {
	schema := new(MysqlBlueprint)
	schema.engine = SchemaDefaultEngine
	schema.name = tableName
	call(schema)
	sql := Assembly(CreateDefaultType, schema) // 拼装成语句
	//fmt.Println(sql)
	stmt, err := ms.GetConn().Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	return err
}

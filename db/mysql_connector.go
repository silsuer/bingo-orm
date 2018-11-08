package db

import (
	"database/sql"
)

type MysqlConnector struct {
	connection   *sql.DB // 数据库连接
	Connector
	name         string // 用户名
	password     string // 密码
	host         string // 主机名
	port         string // 端口
	databaseName string // 数据库名
	charset      string // 字符集
}

// 获取链接
func (m *MysqlConnector) SetConn(config map[string]string) {
	// 开始初始化
	if n, ok := config["db_username"]; ok {
		m.name = n
	} else {
		panic("the mysql connector need db_username")
	}
	if n, ok := config["db_password"]; ok {
		m.password = n
	} else {
		panic("the mysql connector need db_password")
	}
	if n, ok := config["db_host"]; ok {
		m.host = n
	} else {
		panic("the mysql connector need db_host")
	}
	if n, ok := config["db_port"]; ok {
		m.port = n
	} else {
		panic("the mysql connector need db_port")
	}
	if n, ok := config["db_name"]; ok {
		m.databaseName = n
	} else {
		panic("the mysql connector need db_name")
	}
	if n, ok := config["db_charset"]; ok {
		m.charset = n
	} else {
		panic("the mysql connector need db_charset")
	}
	db, err := sql.Open("mysql", m.name+":"+m.password+"@tcp("+m.host+":"+m.port+")/"+m.databaseName+"?charset="+m.charset)
	if err != nil {
		panic(err)
	}
	m.connection = db // 赋值给结构体的属性
}

func (m *MysqlConnector) GetConn() *sql.DB {
	return m.connection
}

// 新建连接器
func NewConnector(config map[string]string) IConnector {
	c := &MysqlConnector{}
	c.SetConn(config)
	return c
}

// 连接器上的Table方法
// 新建一个 Builder 并返回
func (m *MysqlConnector) Table(tableName string) IBuilder {
	return NewMysqlBuilder().SetConn(m).Table(tableName)
}

// 返回一个 数据库/表 构建器
func (m *MysqlConnector) Schema() ISchemaBuilder {
	return NewMysqlSchemaBuilder().SetConn(m)
}

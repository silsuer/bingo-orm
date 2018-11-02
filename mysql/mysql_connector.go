package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/silsuer/bingo-orm"
)

type MysqlConnector struct {
	connection *sql.DB
}

// 获取链接
func (m *MysqlConnector) SetConn(config map[string]string) *sql.DB {
	// Driver.dbConfig = Env.Get("DB_USERNAME") + ":" + Env.Get("DB_PASSWORD") + "@tcp(" + Env.Get("DB_HOST") + ":" + Env.Get("DB_PORT") + ")" + "/" + Env.Get("DB_NAME") + "?" + "charset=" + Env.Get("DB_CHARSET")
	var name string         // 用户名
	var password string     // 密码
	var host string         // 主机地址
	var port string         // 主机端口
	var databaseName string // 数据库名
	var charset string      // 字符集
	// 开始初始化
	if n, ok := config["db_username"]; ok {
		name = n
	} else {
		panic("the mysql connector need db_username")
	}
	if n, ok := config["db_password"]; ok {
		password = n
	} else {
		panic("the mysql connector need db_password")
	}
	if n, ok := config["db_host"]; ok {
		host = n
	} else {
		panic("the mysql connector need db_host")
	}
	if n, ok := config["db_port"]; ok {
		port = n
	} else {
		panic("the mysql connector need db_port")
	}
	if n, ok := config["db_name"]; ok {
		databaseName = n
	} else {
		panic("the mysql connector need db_name")
	}
	if n, ok := config["db_charset"]; ok {
		charset = n
	} else {
		panic("the mysql connector need db_charset")
	}
	db, err := sql.Open("mysql", name+":"+password+"@tcp("+host+":"+port+")/"+databaseName+"?charset="+charset)
	if err != nil {
		panic(err)
	}
	m.connection = db // 赋值给结构体的属性
	return db
}

func (m MysqlConnector) GetConn() *sql.DB {
	//fmt.Println(m)
	return m.connection
}

// 新建连接器
func NewConnector(config map[string]string) bingo_orm.Connector {
	c := &MysqlConnector{}
	c.SetConn(config)
	return c
}

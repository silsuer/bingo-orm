package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// 获取连接器,Table().Get()....
/*
 * 一个方法，获取连接器，
 * 新建一个builder
 * 放入连接器，连接器只负责连接数据库，不负责其他功能
 * 调用 Builder 的 Table 等连贯操作方法即可
 * Builder 实现Builder接口
 */
func GetConnector(config map[string]string) *Connector {
	return nil
}

// 构建一个连接器接口，所有实现这个接口的结构体都可以作为参数传入，每个单独的连接器
type Connector interface {
	SetConn(config map[string]string) *sql.DB // 获取数据库连接器接口,config用于传入参数
	GetConn() *sql.DB                         // 获取数据库连接器接口
}

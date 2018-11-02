package mysql

import (
	"github.com/silsuer/bingo-orm"
	"strings"
	"fmt"
)

type MysqlBuilder struct {
	connector      *bingo_orm.Connector
	bingo_orm.Builder
	bindings       map[string][]string  // 绑定的操作符与列名之间的映射
	columns        []string             // 列名()
	distinct       bool                 // 是否用到了去重查询
	distinctColumn string               // 唯一的列
	from           string               // 表名
	joins          string               // 连接
	wheres         []map[string]string  // where的数组
	groups         string               // 组
	havings        string               // group by 之后的操作
	orders         string               // 排序
	limit          string               // 限制
	offset         string               // 偏移
	unions         []*bingo_orm.Builder //
	unionLimit     string
	unionOffset    string
	unionOrders    string
	lock           bool
}

func (m *MysqlBuilder) SetConnector(db *bingo_orm.Connector) {
	m.connector = db
}

func (m *MysqlBuilder) GetConnector() *bingo_orm.Connector {
	return m.connector
}

// 设置表名
func (m *MysqlBuilder) Table(tableSchema string) bingo_orm.Builder {
	m.from = tableSchema
	return m
}

// 设置distinct
func (m *MysqlBuilder) Distinct(column string) bingo_orm.Builder {
	m.distinct = true
	m.distinctColumn = column
	return m
}

func (m *MysqlBuilder) Get() {
	// 组装Builder
	sql := "select "
	if _, ok := m.bindings["select"]; ok {
		sql += strings.Join(m.bindings["select"], ",")
	} else {
		sql += " * "
	}

	if m.from == "" {
		panic("no table set!")
	} else {
		sql += " from " + m.from + " "
	}

	fmt.Println(sql)

	m.GetConnector()
	//for k, _ := range m.bindings {
	//
	//}
}

package db

type MysqlBuilder struct {
	connector      *Connector
	Builder
	bindings       map[string][]string // 绑定的操作符与列名之间的映射
	columns        []string            // 列名()
	distinct       bool                // 是否用到了去重查询
	distinctColumn string              // 唯一的列
	from           string              // 表名
	joins          string              // 连接
	wheres         []map[string]string // where的数组
	groups         string              // 组
	havings        string              // group by 之后的操作
	orders         string              // 排序
	limit          string              // 限制
	offset         string              // 偏移
	unions         []*Builder          //
	unionLimit     string
	unionOffset    string
	unionOrders    string
	lock           bool
}

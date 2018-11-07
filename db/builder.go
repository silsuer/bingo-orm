package db

type Builder interface {
	GetConnector() *Connector
	SetConnector(connector *Connector)
	Table(tableSchema string) Builder // 设置表名
	Distinct(column string) Builder   // 去重方法
	Get()                             //*sql.Rows
	//Find() *sql.Row
}

// Table("ddd") 得到一个Builder
type Driver struct {
	Builder // Builder中内置一个 构建器，驱动可以调用构建器中的方法
	conn *Connector
}

func (d *Driver) LoadConnector(conn Connector) {
	d.conn = &conn
}

func (d *Driver) LoadBuilder(builder Builder) {
	d.Builder = builder
	builder.SetConnector(d.conn)
}

// 新建驱动
func NewDriver() *Driver {
	return &Driver{}
}

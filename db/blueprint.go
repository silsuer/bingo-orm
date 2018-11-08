package db

type IBlueprint interface {
	Increments(columnName string) IBlueprint
	Nullable() IBlueprint
	Comment(comment string) IBlueprint
	Default(def interface{}) IBlueprint
	String(columnName string) IBlueprint
	StringWithLength(columnName string, length int) IBlueprint
	Integer(columnName string) IBlueprint
	IntegerWithLength(columnName string, length int) IBlueprint
	Index(columns ...string)     // 传入多个列，给这些列添加索引
	UniqueIndex(column string)   // 唯一索引
	FullTextIndex(column string) // 全文索引
	PrimaryKey(column string)    // 主键索引
}

// 表格的插入语句
type Blueprint struct {
}

func (bp *Blueprint) PrimaryKey(column string) {

}

// 添加普通索引，传入一个值代表普通索引，传入多个值代表聚合索引
func (bp *Blueprint) Index(columns ...string) {

}

// 添加聚合索引
func (bp *Blueprint) UniqueIndex(column string) {

}

// 添加全文索引  只对MyISAM表有效
func (bp *Blueprint) FullTextIndex(column string) {

}

func (bp *Blueprint) Increments(columnName string) IBlueprint {
	return bp
}

func (bp *Blueprint) Nullable() IBlueprint {
	return bp
}

func (bp *Blueprint) Comment(comment string) IBlueprint {
	return bp
}

func (bp *Blueprint) Default(def interface{}) IBlueprint {
	return bp
}

func (bp *Blueprint) String(columnName string) IBlueprint {
	return bp
}

func (bp *Blueprint) StringWithLength(columnName string, length int) IBlueprint {
	return bp
}

func (bp *Blueprint) Integer(columnName string) IBlueprint {
	return bp
}

func (bp *Blueprint) IntegerWithLength(columnName string, length int) IBlueprint {
	return bp
}

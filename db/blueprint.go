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
}

// 表格的插入语句
type Blueprint struct {

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

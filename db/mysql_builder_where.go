package db

const (
	Greater = ">"      // 大于
	Equals  = "="      // 等于
	Less    = "<"      // 小于
	Is      = "is"     // is
	IsNot   = "is not" // is not
)

type whereCondition struct {
	columnFunc string      // 传入一个方法名，将会把column作为参数传进去
	columnName string      // 列名
	operator   string      // 比较运算符，默认是 =
	value      interface{} // 值
}

func NewWhereCondition(mb *MysqlBuilder) *whereCondition {
	w := new(whereCondition)
	w.operator = Equals
	mb.wheres = append(mb.wheres, w) // 把where挂载到builder上
	return w
}

// 添加where条件
func (m *MysqlBuilder) Where(condition ...interface{}) IBuilder {
	w := NewWhereCondition(m)
	switch len(condition) {
	case 2:
		if name, ok := condition[0].(string); ok {
			w.columnName = name
		} else {
			panic("the 1st argument in Where function must be string.")
		}
		w.value = condition[1]
	case 3:
		if name, ok := condition[0].(string); ok {
			w.columnName = name
		} else {
			panic("the 1st argument in Where function must be string.")
		}
		if operator, ok := condition[1].(string); ok {
			w.operator = operator
		} else {
			panic("the 2nd argument in Where function must be string.")
		}
		w.value = condition[2]
	default:
		panic("the where condition function need 2 or 3 arguments.")
	}
	return m
}

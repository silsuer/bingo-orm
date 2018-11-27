package db


// 连接条件
type joinOnCondition struct {
	columnFunc string      // 传入一个方法名，将会把column作为参数传进去
	columnName string      // 列名
	operator   string      // 比较运算符，默认是 =
	value      interface{} // 值
}

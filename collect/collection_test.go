package collect

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	// 传入一个数组
	// 传入一个切片
	// 传入一个map
	//Generate([...]string{
	//	"111",
	//	"222",
	//})
	a := make(map[string]int)
	a["111"] = 111
	a["222"] = 222
	b := Generate(a)
	b.Iterator(func(item interface{}) {
		// 集合
	})
}

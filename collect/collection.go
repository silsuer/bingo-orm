package collect

import (
	"sync"
	"reflect"
)

// 仿照laravel集合制作的功能，查询得到的数据
// 对于模型：
// 查询出的单结果，是一个模型对象
// 查询出的多结果，是一个集合对象

// 对于 Table(),查询出的，是一个集合对象，每个对象都是一个map

const (
	TypeArray = iota
	TypeSlice
	TypeMap
)

type Collection struct {
	sync.Mutex                // 读写锁
	items     interface{}     // 集合列表
	current   int             // 当前指针
	length    int             // 长度
	keys      []reflect.Value // 键名
	keyType   reflect.Type    // map的键的类型
	valueType reflect.Type    // map的值得类型
	kind      int             // 集合中的数据类型(数组、切片、map)
	typeOf    reflect.Type    // 反射类型
	valueOf   reflect.Value   // 反射值
}

func Generate(m interface{}) *Collection {
	// 将传入的对象，根据类型生成集合
	c := &Collection{}
	mm := reflect.TypeOf(m)
	mv := reflect.ValueOf(m)
	c.valueOf = mv
	c.typeOf = mm

	switch mm.Kind() {
	case reflect.Slice: // 切片
		c.kind = TypeSlice
	case reflect.Array: // 数组
		c.kind = TypeArray
	case reflect.Map: // map
		c.kind = TypeMap
		c.keys = mv.MapKeys()
		c.keyType = mm.Key()    // map的键的类型
		c.valueType = mm.Elem() // 元素的类型
	default:
		return nil
	}
	// 长度
	c.length = mv.Len()
	return c
}

// 集合迭代器
func (c *Collection) Iterator(call func(item interface{})) {
	call(c.Current())
	for {
		if c.Next() != -1 {
			call(c.Current())
		} else {
			break
		}
	}
}

func (c *Collection) Current() interface{} {
	switch c.kind {
	case TypeArray, TypeSlice:
		return c.valueOf.Index(c.current)
	case TypeMap:
		tmp := make(map[interface{}]interface{})
		tmp[c.keys[c.current]] = c.keys
		return tmp
	}
	return nil
}

// 将指针移动到下一个位置，返回当前位置索引
func (c *Collection) Next() int {
	if c.current+1 < c.length {
		c.current++
		return c.current
	} else {
		return -1
	}
}

// 可用的方法：
// 1. 迭代器
// 2. all
// 3. average
// 4. avg
// 5. chunk
// 6. collapse
// 7. combine
// 8. contains
// 9. containStrict
// 10. count
// 11. diff
// 12. diffKeys
// 13. each
// 14. every
// 15. except
// 16. filter
// 17. first
// 18. flatMap
// 19. flatten
// 20. flip
// 21. forget
// 22. forPage
// 23. get
// 24. groupBy
// 25. has
// 26. implode
// 27. intersect
// 28. isEmpty
// 29. isNotEmpty
// 30. keyBy
// 31. keys
// 32. last
// 33. map
// 34. mapWithKeys
// 35. max
// 36. median
// 37. merge
// 38. min
// 39. mode
// 40. nth
// 41. only
// 42. partition
// 43. pipe
// 44. pluck
// 45. pop
// 46. prepend
// 47. pull
// 48. push
// 49. put
// 50. random
// 51. reduce
// 52. reject
// 53. reverse
// 54. search
// 55. shift
// 56. shuffle
// 57. slice
// 58. sort
// 59. sortBy
// 60. sortByDesc
// 61. splice
// 62. split
// 63. sum
// 64. take
// 65. tap
// 66. toArray
// 67. toJson
// 68. transform
// 69. union
// 70. unique
// 71. uniqueScript
// 72. values
// 73. when
// 74. where
// 75. whereScript
// 76. whereIn
// 77. whereInScript
// 78. whereIn
// 79. whereScript
// 80. whereNotIn
// 81. whereNotInScript
// 82. zip

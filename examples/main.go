package main

import (
	"github.com/silsuer/bingo-orm/db"
	"fmt"
)

func main() {
	// 输入数据
	config := make(map[string]string)
	config["db_username"] = "root"
	config["db_password"] = ""
	config["db_host"] = "127.0.0.1"
	config["db_port"] = "3306"
	config["db_name"] = "test"
	config["db_charset"] = "utf8"
	c := db.NewConnector(config)
	// Table 返回一个 TableBuilder
	//a := c.Table("users")

	// 连接器添加数据
	// Schema() 方法返回一个 SchemaBuilder
	//res, err := c.Schema().CreateDatabaseIfNotExists("test")

	// 创建表
	err := c.Schema().CreateTable("test", func(table db.IBlueprint) {
		table.Increments("id").Comment("自增id")
		table.String("name").Comment("姓名")
		table.Integer("age").Nullable().Comment("年龄")
	})

	fmt.Println(err)
	//fmt.Println(res)
	//query, err := s.Query("select * from users") // Query用作搜索，Exec用来增删改
	//if err != nil {
	//	panic(err)
	//}
	////var name string
	////r.Scan(&name)
	////fmt.Println(name)
	//column, _ := query.Columns()
	//
	//values := make([][]byte, len(column))     //values是每个列的值，这里获取到byte里
	//scans := make([]interface{}, len(column)) //因为每次查询出来的列是不定长的，用len(column)定住当次查询的长度
	//for i := range values { //让每一行数据都填充到[][]byte里面
	//	scans[i] = &values[i]
	//}
	//results := make(map[int]map[string]string) //最后得到的map
	//i := 0
	//for query.Next() { //循环，让游标往下移动
	//	if err := query.Scan(scans...); err != nil { //query.Scan查询出来的不定长值放到scans[i] = &values[i],也就是每行都放在values里
	//		fmt.Println(err)
	//		return
	//	}
	//	row := make(map[string]string) //每行数据
	//	for k, v := range values { //每行数据是放在values里面，现在把它挪到row里
	//		key := column[k]
	//		row[key] = string(v)
	//	}
	//	results[i] = row //装入结果集中
	//	i++
	//}
	//for k, v := range results { //查询出来的数组
	//	fmt.Println(k, v)
	//}

	//fmt.Println(r.Columns())
}

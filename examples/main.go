package main

import (
	"github.com/silsuer/bingo-orm/mysql"
	"fmt"
	"github.com/silsuer/bingo-orm"
)

func main() {
	// 输入数据
	config := make(map[string]string)
	config["db_username"] = "root"
	config["db_password"] = ""
	config["db_host"] = "127.0.0.1"
	config["db_port"] = "3306"
	config["db_name"] = "pengpeng"
	config["db_charset"] = "utf8"

	// 新建连接器
	//c := mysql.MysqlConnector{}
	// 设置连接参数
	//s := c.SetConn(config)

	// 新建连接器
	c := mysql.NewConnector(config)

	// 新建驱动
	bb := bingo_orm.NewDriver()
	// 驱动里加载连接器，通过更换连接器就可以更换数据库，以及读写分离等
	bb.LoadConnector(c) // 加载连接器
	bb.LoadBuilder(&mysql.MysqlBuilder{})
	fmt.Println(bb)
	bb.Table("users").Get()
	// 可以使用连贯操作了
	//c.GetConn().Query()
	s := c.GetConn()
	query, err := s.Query("select * from users") // Query用作搜索，Exec用来增删改
	if err != nil {
		panic(err)
	}
	//var name string
	//r.Scan(&name)
	//fmt.Println(name)
	column, _ := query.Columns()

	values := make([][]byte, len(column))     //values是每个列的值，这里获取到byte里
	scans := make([]interface{}, len(column)) //因为每次查询出来的列是不定长的，用len(column)定住当次查询的长度
	for i := range values { //让每一行数据都填充到[][]byte里面
		scans[i] = &values[i]
	}
	results := make(map[int]map[string]string) //最后得到的map
	i := 0
	for query.Next() { //循环，让游标往下移动
		if err := query.Scan(scans...); err != nil { //query.Scan查询出来的不定长值放到scans[i] = &values[i],也就是每行都放在values里
			fmt.Println(err)
			return
		}
		row := make(map[string]string) //每行数据
		for k, v := range values { //每行数据是放在values里面，现在把它挪到row里
			key := column[k]
			row[key] = string(v)
		}
		results[i] = row //装入结果集中
		i++
	}
	for k, v := range results { //查询出来的数组
		fmt.Println(k, v)
	}

	//fmt.Println(r.Columns())
}

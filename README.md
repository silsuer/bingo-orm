# bingo-orm

bingo-orm 是 [bingo](https://github.com/silsuer/bingo) 框架下的一个子模块，可以快速沟通数据库

特性:
 - [x] 链式操作
 - 模型
 - 数据库迁移
 - 一键开启数据库连接池
 - 数据库假数据填充
    
## 使用方式

 1. 连接数据库
   - 单一连接
     
     ```go
      // 输入数据
	  config := make(map[string]string)
	  config["db_username"] = "root"
	  config["db_password"] = ""
	  config["db_host"] = "127.0.0.1"
	  config["db_port"] = "3306"
	  config["db_name"] = "test"
	  config["db_charset"] = "utf8"
	  c := db.NewConnector(config)  // 传入连接参数，可以得到一个Mysql连接，使用其他数据库则调用其他的新建数据库连接的方法        
     ```
     
   - 使用数据库连接池(待定)
   
 2. 创建数据库
     
   在连接器上调用`Schema()` 方法获得一个 `SchemaBuilder` ，用来对数据库以及数据表结构进行操作  
     
   ```go
      res, err := c.Schema().CreateDatabaseIfNotExists("test")  // 第二个参数是字符集，第三个参数是排序规则
   ```
 
 3. 创建数据表
 
  `SchemaBuilder` 上提供了创建表的方法，在回调函数中进行表结构的定义
  
  ```go
      err := c.Schema().CreateTable("test", func(table db.IBlueprint) {
      		table.Increments("id").Comment("自增id")  // 设置备注与主键
      		table.String("name").Comment("姓名")  
      		table.Integer("age").Nullable().Comment("年龄") // 允许为空
      		
      		// 添加普通索引  _index
            table.Index("user_id")
            // 添加唯一索引  _unique_index
            table.UniqueIndex("user_id")
            // 添加组合索引 
            table.Index("user_id","name")
            // 添加全文索引,只对MyISAM表有效
            table.FullTextIndex("user_id")
     })
  ```
  
## 接下来的任务
  - [x] 创建表可以添加各种类型的字段（float double 等）  
  - [x] 更改表结构
  - [x] 对表进行增删改
  - [x] 重新组织结果集(ToMapList ToStringMapList)
  - [x] 对表进行简单查询
  - [x] mysql事务处理
  - [x] 内外连接，完全连接
  - 实现多态（暂时不实现）
  - 对表查询时指定优先级
  - 对表进行复杂查询(子表等)
  - 添加模型处理（对模型增加观察者，模型与db进行关联，底层使用db进行操作）
  - 读写分离
  - 连接池
  - 整理说明文档
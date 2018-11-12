# bingo-orm

bingo-orm 是 [bingo](https://github.com/silsuer/bingo) 框架下的一个子模块，可以快速沟通数据库

特性:
 - 链式操作
 - 模型
 - 数据库迁移
 - 一键开启数据库连接池
 - 数据库假数据填充
 
 # 流程
 
 1. 连接
   
    创建一个工厂
    
    用工厂创建一个连接器
    
# 使用方式

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
     
   - 使用数据库连接池
   
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
  - 实现多态（暂时不实现）
  - [x] 更改表结构
  - 对表进行增删改
  - 重新组织结果集
  - 对表进行查询(内连接 外连接 子表)
  - [x] mysql事务处理
  - 添加模型处理（对模型增加观察者，模型与db进行关联，底层使用db进行操作）
  - 整理说明文档
## 知识点
   
   mysql的各种类型
   mysql的索引
   mysql的各种连接
   
   https://blog.csdn.net/gxy_2016/article/details/53436865  日期类型的区别
   https://blog.csdn.net/MinjerZhang/article/details/78137795 mysql 地理位置
   http://www.cnblogs.com/cnsanshao/p/3326648.html  数字和ip地址相互转换
   https://blog.csdn.net/zzc_zcc/article/details/78836505 类型
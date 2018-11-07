# bingo-orm

bingo-orm 是 [bingo](https://github.com/silsuer/bingo) 框架下的一个子模块，可以快速沟通数据库

特性:
 - 链式操作
 - 模型
 - 数据库迁移
 - 一键开启数据库连接池
 - 数据库假数据填充
 
 // 两种操作，一种是直接操作Table，结果指针
 // 另一种返回对应的模型
 
 对每一张表都要生成一个模型，
 模型要对应一个builder，最终返回数据
 
 // 定义一个接口
 // MYSQL实现这个接口，重新实现各种方法，接口接受一个实现 connector的方法，放入connection
 
 
 # 流程
 
 1. 连接
   
    创建一个工厂
    
    用工厂创建一个连接器
    
    
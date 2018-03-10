# gbdb
   GbDB   折腾数据库玩具

### 挖坑相关
18年3月8日正式开始挖坑,辞去了北京的实习工作，回家可能呆1个月左右。
### 文档相关

以下是基本概念，具体详知请看《MySQL 技术内幕》《MySQL运维内参》, innodb储存文件观察 py_innodb_page_info .以下是摸索阶段的入门概念,相比《数据库系统实现》更
为亲民一些.

[索引查询原理](http://blog.jobbole.com/24006/)

[B树相关实现](https://www.cnblogs.com/vincently/p/4526560.html)

[覆盖索引&聚集索引&非聚集索引概念](https://www.cnblogs.com/aspwebchh/p/6652855.html)


[理解B+树算法和Innodb索引](https://www.cnblogs.com/huqiang/p/5604722.html)

[innodb 页概念](https://segmentfault.com/a/1190000008545713)


### 挖坑阶段

#### 1. 物理文件储存逻辑
   规划,看到简单的实现基本的都是以0号页储存表元信息,采用页的形式来组织整个储存结构。gbdb还是使用类似于innodb采用系统表空间的方法构建索引和表文件结构记录等
   这一阶段主要实现物理文件结构的划分和实现。
   参考:
   
[innodb空间管理页结构](https://blog.jcole.us/2013/01/04/page-management-in-innodb-space-files/)
[innodb引擎页结构](https://dev.mysql.com/doc/internals/en/innodb-page-overview.html)   
[innodb储存结构](https://www.kancloud.cn/digest/innodb-zerok/195090)

s
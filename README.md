# gbdb
   GbDB   折腾数据库玩具

### 挖坑相关


### 文档相关

以下是基本概念，具体详知请看《MySQL 技术内幕》《MySQL运维内参》, innodb储存文件观察 py_innodb_page_info .以下是摸索阶段的入门概念,相比《数据库系统实现》更
为亲民一些.


### 挖坑阶段
太菜了，脑汁里一团乱麻 = = ！
#### 1. 物理文件储存逻辑
   规划,看到简单的实现基本的都是以0号页储存表元信息,采用页的形式来组织整个储存结构。gbdb还是使用类似于innodb采用系统表空间的方法构建索引和表文件结构记录等
   这一阶段主要实现物理文件结构的划分和实现。
####  2018.4.2
   space file management 页面结构
####  2018.10.12
   - pass: 已知 稍微完成 cacahe parser 相关实现

   - 设计 space -> table  引入syscache

   - 目前工作 重新设计space file management 为索引树创建管理提供支持

当前模块设计:
   - cache 页缓存池 页锁bufferpage  全局系统字典表syscache
   - dm 行记录
   - im 索引
   - log redo/undo 记录
   - parser sql 解析
   - server server入口点
   - spaceManage 表空间管理器
   - tbm 表管理 tfm 表文件管理
   - tm 事物管理器



### note 入口
[笔记](https://github.com/fangker/gbdb/blob/master/note.md)
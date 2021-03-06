package _doc

/*
ParseFile
数据库文件format

# 1
frm
-Header

|2|4|6|8|10|12|14|16|18|……|64

[1:2]   frm类型
[3:4]   储存引擎
[5:6]   跳过x字节为字段信息存储位置  default: 2048
[7:8]   从指定定位置读取字节数
[9:10]   key存储位置
[11:12]  key 信息长度
[13:14] 一个key长度
[15:16] 读取默认1行数据长调度
[17:18] 表字符集
[19]    行类型


-Key

total [6]
[1]   key数量
[2]	  字段数量
[3]   key后y位存放key-name 描述信息
[4]	  key描述信息后x位为comment信息
[5]

key
[1:2] 预留
[3]   有几个key part
[4]   索引类型
[5]
[6]
[7]

key part

[1:2] 第为表中x个字段
[3:4] 数据行偏移量 compact 0 填充
[5:6] 字段长度
[7:8]


field header
total 300
[1:4] 表注释长度
[5:8] 字段名拼接长度
表注释内容
所有字段名拼接长度
拼接内容

field

[1:2]
[3:4] 字段长度
[5:6] 字段数据行偏移量
[7:8] 字段字符集
[9]   类型






*/
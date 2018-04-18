package db_go

//数据库连接结构
type DB_Connection struct {
	//要在其他结构调用结构体，需要属性大写 否则报错：implicit assignment of unexported field
	DB_Connection_userName string //数据库用户名
	DB_Connection_password string //数据库用户密码
	DB_Connection_host     string //数据库主机地址
	DB_Connection_port     string //数据库地址端口号
}

package db_go

import (
	"database/sql"
	"fmt"
)

//数据库连接结构
type DB_Database struct {
	//要在其他结构调用结构体，需要属性大写 否则报错：implicit assignment of unexported field
	DB_Database_name       string //数据库用户名
	DB_Database_collate    string //数据库默认字符集
	DB_Database_character  string //数据库默认排序规则
	DB_Database_connection DB_Connection
}

//创建数据库
func (database DB_Database) Create(db *sql.DB) {
	result, err := db.Exec(splitSQL_createDatabase(database))
	fmt.Print(splitSQL_createDatabase(database))
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(result)
}

//拼接创建数据库SQL语句
func splitSQL_createDatabase(database DB_Database) string {
	sqlstr := "create database " +
		database.DB_Database_name +
		" default character set " +
		database.DB_Database_character +
		" collate " +
		database.DB_Database_collate
	return sqlstr
}

//连接数据库信息成链接
func LinkURL(database DB_Database) string {
	urlstr := database.DB_Database_connection.DB_Connection_userName +
		":" +
		database.DB_Database_connection.DB_Connection_password +
		"@tcp(" + database.DB_Database_connection.DB_Connection_host +
		":" +
		database.DB_Database_connection.DB_Connection_port +
		")/" +
		database.DB_Database_name

	fmt.Println("link is " + urlstr)

	return urlstr
}

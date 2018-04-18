package db_go

import (
	"database/sql"
	"fmt"
)

//目标表
type DB_Table struct {
	DB_Table_name     string
	DB_Table_database DB_Database
}

//创建数据库
func (table DB_Table) Create(db *sql.DB) {
	result, err := db.Exec(splitSQL_createTable(table))
	fmt.Print(splitSQL_createTable(table))
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(result)
}

//拼接创建数据库SQL语句
func splitSQL_createTable(table DB_Table) string {
	sqlstr := "create table " +
		table.DB_Table_name +
		" ( " +

		" ) "
	return sqlstr
}

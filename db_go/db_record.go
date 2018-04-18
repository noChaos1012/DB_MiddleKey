package db_go

import (
	"database/sql"
	"fmt"
	"log"
)

type DataType struct {
}

const (
	T_int    = "int"
	T_string = "string"
)

//目标表
type DB_Record struct {
	DB_Record_key   string
	DB_Record_value string
	DB_Record_Type  string
	DB_Record_table DB_Table
}

//新增
func InsertSingle(db *sql.DB, record DB_Record) {
	result, err := db.Exec("INSERT INTO " +
		record.DB_Record_table.DB_Table_name +
		"(" +
		record.DB_Record_key +
		") VALUES" +
		record.DB_Record_value)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(result)
}

//查询
func (record DB_Record) Query(db *sql.DB) {
	rows, err := db.Query("SELECT " +
		record.DB_Record_key +
		"FROM " +
		record.DB_Record_table.DB_Table_name)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var name string
		var id int
		if err := rows.Scan(&name, &id); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d : %s \n", id, name)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

//修改
func (toRecord DB_Record) UpdateSingle(db *sql.DB, fromRecord DB_Record) {
	result, err := db.Exec("update " +
		fromRecord.DB_Record_table.DB_Table_name +
		" set " +
		toRecord.DB_Record_key +
		"= " +
		toRecord.DB_Record_value +
		" where " +
		fromRecord.DB_Record_key +
		"=" +
		fromRecord.DB_Record_value)

	if err != nil {
		panic(err)
		return
	}
	fmt.Println(result)
}

//删除
func (record DB_Record) DeleteSingle(db *sql.DB) {
	//stm, _ := db.Prepare("DELETE FROM t_test WHERE id=?")
	//stm.Exec(1)
	//stm.Close()

	result, err := db.Exec("DELETE FROM " +
		record.DB_Record_table.DB_Table_name +
		" WHERE " +
		record.DB_Record_key +
		"=" +
		record.DB_Record_value)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(result)
}

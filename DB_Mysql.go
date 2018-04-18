package main

import (
	"DB_MiddleKey/db_init"
	_ "github.com/Go-SQL-Driver/MySQL"
)

func main() {

	//fmt.Print(baseConfig.LOCALDATABASECONFIG.USERNAME)
	db_init.InitStruct()

	//connection := db_go.DBConnection{
	//	"dbm",
	//	"waschilddbm",
	//	"192.168.1.21",
	//	"3306",
	//	"mysql",
	//}
	//table := db_go.TagertTable{
	//	"t_test",
	//}

	//table.InsertName(db, "aaaccc")
	//table.Update(db, 2, "waschild")
	//table.Delete(db, 10)
	//table.Query(db)
}

package db_init

import (
	"DB_MiddleKey/db_go"
	"database/sql"
	"encoding/json"
	"fmt"
	io "io/ioutil"
)

type BaseConfigStruct struct {
	LOCAL_DB_CONFIG LocalDatabaseConfigStruct
}

type LocalDatabaseConfigStruct struct {
	DB_USERNAME  string
	DB_PASSWORD  string
	DB_COLLATE   string
	DB_CHARACTER string
	DB_HOST      string
	DB_PORT      string
}

const default_databaseName = "Middlekey_BaseDB" //基础库

var baseconfig BaseConfigStruct

func InitStruct() {
	baseconfig = LoadBaseConfig("baseConfig.json")
	fmt.Printf("baseconfig is %+v", baseconfig)

	defaultDB := db_go.DB_Database{
		DB_Database_name: "mysql",
		DB_Database_connection: db_go.DB_Connection{
			DB_Connection_userName: baseconfig.LOCAL_DB_CONFIG.DB_USERNAME,
			DB_Connection_password: baseconfig.LOCAL_DB_CONFIG.DB_PASSWORD,
			DB_Connection_host:     baseconfig.LOCAL_DB_CONFIG.DB_HOST,
			DB_Connection_port:     baseconfig.LOCAL_DB_CONFIG.DB_PORT,
		},
	}

	fmt.Printf("baseconfig.LOCAL_DB_CONFIG.DB_PORT is %s", baseconfig.LOCAL_DB_CONFIG.DB_PORT)

	db := CreateConnection(defaultDB)
	base_database := db_go.DB_Database{
		DB_Database_name:       default_databaseName,
		DB_Database_collate:    baseconfig.LOCAL_DB_CONFIG.DB_COLLATE,
		DB_Database_character:  baseconfig.LOCAL_DB_CONFIG.DB_CHARACTER,
		DB_Database_connection: defaultDB.DB_Database_connection,
	}

	//创建基础数据库
	base_database.Create(db)
	defer db.Close()

}

//连接本地数据库
func CreateConnection(defaultDB db_go.DB_Database) *sql.DB {
	db, err := sql.Open("mysql",
		db_go.LinkURL(defaultDB))
	if err != nil {
		panic(err)
		return nil
	}
	return db
}

//读取基础配置文件
func LoadBaseConfig(filename string) BaseConfigStruct {

	config := BaseConfigStruct{}

	data, err := io.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(data), &config)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%s %+v \n", []byte(data), config)
	fmt.Printf("loadbaseconfig log  %s  %s \n", config.LOCAL_DB_CONFIG.DB_USERNAME, config.LOCAL_DB_CONFIG.DB_CHARACTER)
	return config
}

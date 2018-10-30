package database

import (
	"blog_gin/pkg/setting"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var SqlDB *sql.DB

func init() {
	var err error
	
	DB_USER := setting.MysqlSetting.User
	DB_PASS := setting.MysqlSetting.Pwd
	DB_HOST := setting.MysqlSetting.Host
	DB_PORT := setting.MysqlSetting.Port
	DB_NAME := setting.MysqlSetting.DbName
	
	SqlDB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",DB_USER, DB_PASS, DB_HOST,DB_PORT, DB_NAME))

	if err != nil {
		log.Fatal(err.Error())
	}
	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}

package database

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"database/sql"
	"fmt"
	"blog_gin/pkg/setting"
)

//var SqlDB *xorm.Engine
var SqlDB *sql.DB

func init() {
	var err error

	DB_USER := setting.MysqlSetting.USER
	DB_PASS := setting.MysqlSetting.PASSWORD
	DB_HOST := setting.MysqlSetting.HOST
	DB_PORT := setting.MysqlSetting.PORT
	DB_NAME := setting.MysqlSetting.NAME
	//DB_TYPE := setting.MysqlSetting.TYPE
	//SqlDB, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/person?parseTime=true")
	SqlDB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",DB_USER, DB_PASS, DB_HOST,DB_PORT, DB_NAME))
	//SqlDB, err = xorm.NewEngine("mysql", fmt.Sprintf("%s:%v@tcp(%s:%d)/%s?charset=utf8", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME))
	if err != nil {
		log.Fatal(err.Error())
	}
	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}

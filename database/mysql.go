package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

var Engine *xorm.Engine

func init() {
	var err error

	Engine,err := xorm.NewEngine("mysql","root:Test123...@tcp(119.3.11.44:3306)/blog?parseTime=true")

	if err != nil {
		fmt.Println("新建引擎", err)
		return
	}

	if err := Engine.Ping(); err != nil {
		fmt.Println(err)
		return
	}
	Engine.SetTableMapper(core.SameMapper{})
	Engine.ShowSQL(true)
	Engine.SetMaxOpenConns(5)

	/*
	DB_USER := setting.MysqlSetting.USER
	DB_PASS := setting.MysqlSetting.PASSWORD
	DB_HOST := setting.MysqlSetting.HOST
	DB_PORT := setting.MysqlSetting.PORT
	DB_NAME := setting.MysqlSetting.NAME
	DB_TYPE := setting.MysqlSetting.TYPE
	SqlDB, err := xorm.NewEngine("mysql", fmt.Sprintf("%s:%v@tcp(%s:%d)/%s?charset=utf8", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME))
	if err != nil {
		log.Fatal(err.Error())
	}
	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}*/
}

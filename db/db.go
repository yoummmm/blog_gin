package db

import (
	"blog_gin/pkg/setting"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/lunny/log"
)

//gloabl db
var engin *xorm.Engine

func init(){
	var err error
	DB_USER := setting.MysqlSetting.USER
	DB_PASS := setting.MysqlSetting.PASSWORD
	DB_HOST := setting.MysqlSetting.HOST
	/*DB_USER := setting.MysqlSetting.USER
	DB_PASS := setting.MysqlSetting.PASSWORD
	DB_HOST := setting.MysqlSetting.HOST
	DB_PORT := setting.MysqlSetting.PORT
	DB_NAME := setting.MysqlSetting.NAME
	DB_TYPE := setting.MysqlSetting.TYPE
	SqlDB, err := xorm.NewEngine("mysql", fmt.Sprintf("%s:%v@tcp(%s:%d)/%s?charset=utf8", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME))*/
	source := fmt.Sprintf("%s:%s@%s", DB_USER, DB_PASS, DB_HOST)
	engin, err = xorm.NewEngine("mysql", source)
	
	if err != nil {
		log.Fatalf("db error: %#v\n", err.Error())
	}

	err = engin.Ping()
	if err != nil {
		log.Fatalf("db connect error: %#v\n", err.Error())
	}
	
}

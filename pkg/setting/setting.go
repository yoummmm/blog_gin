package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	cfg *ini.File

	RunMode string

	HttpPort     int
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration

	PageSize int
)

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

type Mysql struct {
	HOST         string
	PORT         int
	PASSWORD     string
	TYPE         string
	USER         string
	NAME         string
	TABLE_PREFIX string
}

var RedisSetting = &Redis{}
var MysqlSetting = &Mysql{}

func init() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatal("gen配置文件打开错误 'conf/app.ini': %v", err)
	}
	mapTo("database", MysqlSetting)
	mapTo("redis", RedisSetting)
	LoadBase()
	LoadServer()
	LoadApp()
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}

func LoadBase() {
	RunMode = cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	servers := cfg.Section("server")
	HttpPort = servers.Key("HTTP_PORT").MustInt(8001)

	ReadTimeOut = time.Duration(servers.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeOut = time.Duration(servers.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	server, err := cfg.GetSection("app")
	if err != nil {
		log.Fatal("配置文件打开错误 'conf/app.ini': %v", err)
	}
	PageSize = server.Key("PAGE_SIZE").MustInt(12)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err:%v", err)
	}
}

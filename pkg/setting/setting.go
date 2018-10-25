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

func init() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatal("gen配置文件打开错误 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
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

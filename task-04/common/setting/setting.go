package setting

import (
	"github.com/go-ini/ini"
	"log"
)

var DatabaseSetting = &Database{}
var LoggerSetting = &Logger{}
var JWTSetting = &Jwt{}
var ServerSetting = &Server{}

var conf *ini.File

// Load 加载配置文件
func Load() {
	var err error
	conf, err = ini.Load("app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'app.ini': %v", err)
	}
	loadSelection("database", DatabaseSetting)
	loadSelection("logger", LoggerSetting)
	loadSelection("jwt", JWTSetting)
	loadSelection("server", ServerSetting)
}

func loadSelection(selection string, v interface{}) {
	err := conf.Section(selection).MapTo(v)
	if err != nil {
		log.Fatalf("Fail to map to DatabaseSetting: %v", err)
	}
}

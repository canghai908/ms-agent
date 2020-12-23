package setting

import (
	"github.com/go-ini/ini"
	log "github.com/sirupsen/logrus"
)

var cfg *ini.File

type App struct {
	Debug       string
	TenantID    string
	LogSavePath string
	Host        string
	Token       string
}

var AppSetting = &App{}

func Setup() {
	var err error
	cfg, err = ini.Load("/etc/ms-agent/app.ini")
	if err != nil {
		log.Errorln("Fail to parse 'app.ini': %v", err)
		return
	}
	mapTo("app", AppSetting)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Errorln("Cfg.MapTo RedisSetting err: %v", err)
		return
	}
}

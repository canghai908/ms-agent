package setting

import (
	"github.com/go-ini/ini"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
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
	Mpath, err := os.Executable()
	if err != nil {
		log.Errorln(err)
	}
	cfg, err = ini.Load(path.Dir(Mpath) + "/app.ini")
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

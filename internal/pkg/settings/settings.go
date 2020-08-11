package settings

import (
	"github.com/go-ini/ini"
	"log"
)

type Database struct {
	Type string
	User string
	Password string
	Host string
	Name string
}

var DBSetting = &Database{}

func InitSetting() {
	conf, err := ini.Load("internal/conf/conf.ini")
	if err != nil {
		log.Fatal("load conf.ini failed: ", err)
	}
	err = conf.Section("database").MapTo(DBSetting)
	if err != nil {
		log.Fatal("read database conf failed: ", err)
	}
}
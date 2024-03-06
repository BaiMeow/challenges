package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/jpillora/overseer/fetcher"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

var Conf config
var Fetch fetcher.HTTP

type config struct {
	NoAdminLogin bool
	DBUser       string
	DBPasswd     string
	DBHost       string
	DBPort       string
	AutoUpdate   bool
	UpdateUrl    string
	UpdateTime   time.Duration
}

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
	UpdateConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		UpdateConfig()
	})
	viper.WatchConfig()
}

func UpdateConfig() {
	Conf.DBUser = viper.GetString("database.user")
	Conf.DBPasswd = viper.GetString("database.password")
	Conf.DBHost = viper.GetString("database.host")
	Conf.DBPort = viper.GetString("database.port")
	Conf.NoAdminLogin = viper.GetBool("server.noAdminLogin")
	Conf.AutoUpdate = viper.GetBool("update.enabled")
	Fetch = fetcher.HTTP{
		URL:      viper.GetString("update.url"),
		Interval: viper.GetDuration("update.interval") * time.Second,
	}
	Fetch.Init()
	log.Println("config updated")
}

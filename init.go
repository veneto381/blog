package main

import (
	"log"
	"strings"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	domain string
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("配置文件读取失败\n" + err.Error())
	}
}

func init() {
	if viper.GetBool("database.mysql.enable") {
		dsn := strings.Join([]string{viper.GetString("database.mysql.username"), ":", viper.GetString("database.mysql.password"), "@tcp(", viper.GetString("database.mysql.ip"), ":", viper.GetString("database.mysql.port"), ")/", viper.GetString("database.mysql.dbname"), "?charset=utf8mb4&parseTime=True&loc=Local"}, "")
		log.Println(dsn)
		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("数据库连接失败\n" + err.Error())
		}
	}
}

func init() {
	domain = viper.GetString("domain")
}

package main

import (
	"log"
	"strings"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
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
		dsn := strings.Join([]string{viper.GetString("mysql.username"), ":", viper.GetString("mysql.password"), "@tcp(", viper.GetString("mysql.ip"), ":", viper.GetString("mysql.port"), ")/", viper.GetString("mysql.dbname"), "?charset=utf8mb4&parseTime=True&loc=Local"}, "")
		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("数据库连接失败\n" + err.Error())
		}
	}
}

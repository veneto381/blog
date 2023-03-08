package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	g := gin.Default()

	user := g.Group("/user")
	{
		user.GET("/info/:name", GetUserInfo)
		user.GET("/detail/:name", CheckUser, GetUserDetail)
	}
	g.POST("/login", Login)
	g.POST("/register", Register)
	g.Run(viper.GetString("listenAddr"))
}

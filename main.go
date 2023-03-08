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
	article := g.Group("/article")
	{
		article.GET("/view/:id", GetArticleById)
		article.GET("/titles", GetArticleTitles)
		article.GET("/review", CheckUser, ReviewArticlesList)
		article.POST("", CheckUser, PostArticle)
		article.POST("/review", CheckUser, ReviewArticle)
	}
	g.POST("/login", Login)
	g.POST("/register", Register)
	g.Run(viper.GetString("listenAddr"))
}

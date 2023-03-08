package main

import "github.com/gin-gonic/gin"

func main() {
	g := gin.Default()

	user := g.Group("/user")
	{
		user.GET("/info/:name", GetUserInfo)
	}
}

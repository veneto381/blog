package main

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUserInfo(c *gin.Context) {
	var user User
	name := c.Param("name")
	if errors.Is(db.Where("user_name = ?", name).First(&user).Error, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{
			"errors": []gin.H{
				{
					"code":  USER_NOT_FOUND,
					"title": "user not found",
				},
			},
		})
		return
	}
	c.JSON(200, gin.H{
		"type":     "userinfo",
		"id":       user.Id,
		"username": user.UserName,
		"nickname": user.NickName,
	})
}

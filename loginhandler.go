package main

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.JSON(401, gin.H{
			"errors": []gin.H{
				{
					"code":  BAD_PARAMETERS,
					"title": "bad body",
				},
			},
		})
		return
	}

	var user User
	if errors.Is(db.Where("user_name = ?", username).First(&user).Error, gorm.ErrRecordNotFound) {
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
	if !ComparePassword(password, string(user.PassWord)) {
		c.JSON(200, gin.H{
			"errors": []gin.H{
				{
					"code":  WRONG_PASSWORD,
					"title": "密码错误",
				},
			},
		})
		return
	}

	token, err := GenerateToken(user.Id, user.Type, user.UserName)
	if err != nil {
		c.JSON(500, gin.H{
			"errors": []gin.H{
				{
					"code":  INTERNAL_ERROR,
					"title": err,
				},
			},
		})
		return
	}
	c.SetCookie("token", token, 3600, "/", domain, false, true)
	c.JSON(200, gin.H{
		"type":     "logininfo",
		"status":   "ok",
		"redirect": "/",
	})
}

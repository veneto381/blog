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

func GetUserDetail(c *gin.Context) {
	name := c.Param("name")
	claims, ok := c.Get("claim")
	if !ok {
		c.JSON(403, gin.H{
			"errors": []gin.H{
				{
					"code":  UNAUTHORIZED,
					"title": "unauthorized",
				},
			},
		})
		return
	}
	claim := claims.(*Claims)
	if claim.Username != name {
		c.JSON(403, gin.H{
			"errors": []gin.H{
				{
					"code":  UNAUTHORIZED,
					"title": "unauthorized",
				},
			},
		})
		return
	}
	var user User
	username := c.Param("name")
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
	c.JSON(200, gin.H{
		"type":      "userdetail",
		"id":        user.Id,
		"username":  user.UserName,
		"nickname":  user.NickName,
		"email":     user.Email,
		"birthday":  user.Birthday,
		"telephone": user.Telephone,
		"createAt":  user.CreatedAt,
	})
}

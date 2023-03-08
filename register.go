package main

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(c *gin.Context) {
	username := c.PostForm("username")
	var user User
	if !errors.Is(db.Where("user_name = ?", username).First(&user).Error, gorm.ErrRecordNotFound) {
		c.JSON(401, gin.H{
			"errors": []gin.H{
				{
					"code":  BAD_PARAMETERS,
					"title": "该用户已注册",
				},
			},
		})
		return
	}

	password := c.PostForm("password")
	nickname := c.PostForm("nickname")
	email := c.PostForm("email")
	telephone := c.PostForm("telephone")
	birthdayString := c.PostForm("birthday")
	var birthday *time.Time = nil
	if birthdayString != "" {
		b, err := time.ParseInLocation("2006-01-02", birthdayString, time.Local)
		if err != nil {
			c.JSON(401, gin.H{
				"errors": []gin.H{
					{
						"code":  BAD_PARAMETERS,
						"title": "生日格式不正确",
					},
				},
			})
			return
		}
		birthday = &b
	}
	hashedPassword, err := HashAndSort(password)
	if err != nil {
		c.JSON(500, gin.H{
			"errors": []gin.H{
				{
					"code":  INTERNAL_ERROR,
					"title": "内部错误",
				},
			},
		})
	}
	db.Create(&User{
		UserName:  username,
		PassWord:  hashedPassword,
		NickName:  nickname,
		Email:     email,
		Telephone: telephone,
		Birthday:  birthday,
	})
	c.JSON(200, gin.H{
		"type":   "registerInfo",
		"stauts": "ok",
	})
}

package main

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetArticleById(c *gin.Context) {
	id := c.Param("id")
	var article Article
	if errors.Is(db.Where("id = ?", id).First(&article).Error, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{
			"errors": []gin.H{
				{
					"code":  ARTICLE_NOT_FOUND,
					"title": "article not found",
				},
			},
		})
		return
	}
	c.JSON(200, gin.H{
		"type": "article",
		"data": article,
	})
}

func GetArticleTitles(c *gin.Context) {
	DB := db
	titles, err := strconv.Atoi(c.DefaultQuery("number", "10"))
	if err != nil {
		c.JSON(401, gin.H{
			"errors": []gin.H{
				{
					"code":  BAD_PARAMETERS,
					"title": "bad query",
				},
			},
		})
	}
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(401, gin.H{
			"errors": []gin.H{
				{
					"code":  BAD_PARAMETERS,
					"title": "bad query",
				},
			},
		})
	}
	articleType, ok := c.GetQuery("type")
	if !ok {
		var temp []Article
		db.Order("created_at desc").Offset((page - 1) * titles).Limit(titles).Find(&temp)
		c.JSON(200, gin.H{
			"type": "articles",
			"data": temp,
		})
		return
	}

	var type1 Type
	if errors.Is(db.Select("id").Where("name = ?", articleType).First(&type1).Error, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{
			"errors": []gin.H{
				{
					"code":  TYPE_NOT_FOUND,
					"title": "未知的类型",
				},
			},
		})
		return
	}
	DB.Where("id = ?", type1.Id)
	var types2 []Type
	db.Select("id").Where("father = ?", type1.Id).Find(&types2)
	for _, v := range types2 {
		DB.Where("id = ?", v.Id)
	}

	articles := []Article{}

	DB.Order("created_at desc").Offset((page - 1) * titles).Limit(titles).Find(&articles)
	c.JSON(200, gin.H{
		"type": "articles",
		"data": articles,
	})
}

func PostArticle(c *gin.Context) {
	cl, ok := c.Get("claim")
	if !ok {
		c.JSON(403, gin.H{
			"errors": []gin.H{
				{
					"code":  UNAUTHORIZED,
					"title": "未登陆",
				},
			},
		})
		return
	}
	claim := cl.(*Claims)
	title := c.PostForm("title")
	text := c.PostForm("text")
	articleType, err := strconv.Atoi(c.PostForm("type"))
	if err != nil {
		c.JSON(401, gin.H{
			"errors": []gin.H{
				{
					"code":  BAD_PARAMETERS,
					"title": "文章类型不正确",
				},
			},
		})
		return
	}
	lableString := c.PostForm("lable")
	var lable int
	if lableString == "" {
		lable = 0
	} else {
		lable, err = strconv.Atoi(lableString)
		if err != nil {
			c.JSON(401, gin.H{
				"errors": []gin.H{
					{
						"code":  BAD_PARAMETERS,
						"title": "文章类型不正确",
					},
				},
			})
			return
		}
	}

	var article = Article{
		Title:  title,
		Type:   articleType,
		Author: claim.UserId,
	}
	if lable != 0 {
		article.Lable = lable
	}
	db.Create(&article)
	db.Create(&ArticleContent{
		Id:      article.Id,
		Content: text,
	})
	c.JSON(200, gin.H{
		"type":   "info",
		"status": "ok",
	})
}

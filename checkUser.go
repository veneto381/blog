package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("test")

type Claims struct {
	UserId   int32  `json:"userid"`
	UserType int32  `json:"usertype"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(userid, usertype int32, username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claim := Claims{
		UserId:   userid,
		UserType: usertype,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: expireTime},
			Issuer:    "admin",
		},
	}
	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := tokenClaim.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaim, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaim != nil {
		if claim, ok := tokenClaim.Claims.(*Claims); ok && tokenClaim.Valid {
			return claim, nil
		}
	}
	return nil, err
}

func CheckUser(c *gin.Context) {
	token, err := c.Request.Cookie("token")
	if err != nil {
		c.JSON(403, gin.H{
			"errors": []gin.H{
				{
					"code":     UNAUTHORIZED,
					"title":    "未授权",
					"redirect": "/login",
				},
			},
		})
		c.Abort()
		return
	}
	claim, err := ParseToken(token.Value)
	if err != nil {
		c.JSON(403, gin.H{
			"errors": []gin.H{
				{
					"code":     UNAUTHORIZED,
					"title":    "未授权",
					"redirect": "/login",
				},
			},
		})
		c.Abort()
		return
	}
	if claim.ExpiresAt.Time.Unix() < time.Now().Unix() {
		c.JSON(403, gin.H{
			"errors": []gin.H{
				{
					"code":     UNAUTHORIZED,
					"title":    "未授权",
					"redirect": "/login",
				},
			},
		})
		c.Abort()
		return
	}
	if claim.ExpiresAt.Time.Unix() > time.Now().Add(5*time.Minute).Unix() {
		newToken, err := GenerateToken(claim.UserId, claim.UserType, claim.Username)
		if err != nil {
			c.JSON(400, gin.H{
				"status": "error",
			})
			c.Abort()
			return
		}
		c.Set("updateToken", true)
		c.Set("token", newToken)
	}
	c.Set("claim", claim)

	c.Next()
}

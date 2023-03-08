package main

import "time"

type User struct {
	Id        int32   `gorm:"not null; auto_increment"`
	Type      int32   `gorm:"default:101"` //1站长 2管理员 101用户
	UserName  string  `gorm:"not null"`
	PassWord  []byte  `gorm:"not null; type: char(60)"`
	Coin      float64 `gorm:"default: 0"`
	NickName  string
	Email     string
	Telephone string
	Birthday  *time.Time `gorm:"type: date; default: null"`
	CreatedAt time.Time
}

type Article struct {
	Id            int32  `gorm:"not null; auto_increment"`
	Title         string `gorm:"not null; type: varchar(1024); not null"`
	Summary       string `gorm:"type: varchar(1024)"`
	Lable         int    `gorm:"default: null"`
	Type          int
	Coin          float64 `gorm:"default: 0"`
	CommentNumber int
	LikeNumber    int
	Views         int
	CreatedAt     time.Time
	OnTop         bool `gorm:"default: false"`
	Author        int32
}

type ArticleContent struct {
	Id      int32 `gorm:"not null"`
	Content string
}

type Type struct {
	Id        int32  `gorm:"not null; auto_increment"`
	Name      string `gorm:"not null"`
	Summary   string
	Father    int32
	CreatedAt time.Time
}

type Lable struct {
	Id        int32  `gorm:"not null; auto_increment"`
	Name      string `gorm:"not null"`
	Summary   string
	CreatedAt time.Time
}

type Commentary struct {
	Id         int32 `gorm:"not null; auto_increment"`
	UserId     int32
	Father     int32
	ArticleId  int32
	CreatedAt  time.Time
	Content    string
	LikeNumber int
}

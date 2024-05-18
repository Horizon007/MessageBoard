package model

import "gorm.io/gorm"

func (Message) TableName() string {
	return "message"
}

type Message struct {
	Id         int    `gorm:"primary_key,type:INT;not null;AUTO_INCREMENT"`
	User_Id    int    `json:"User_Id"  binding:"required"`
	Content    string `json:"Content"  binding:"required"`
	Version    int    `gorm:"default:0"`
	gorm.Model        //这个是gorm的内置结构体，包含了ID、CreatedAt、UpdatedAt、DeletedAt这四个字段
}

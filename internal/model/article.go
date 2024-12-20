package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Preview string `json:"preview" binding:"required"`
	Likes   int    `gorm:"default:0" json:"likes"`
}

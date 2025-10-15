package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `gorm:"not null" json:"content"`
	UserID  uint   `json:"user_id" example:"1"`
	User    User
	PostID  uint `json:"post_id" example:"1"`
	Post    Post
}

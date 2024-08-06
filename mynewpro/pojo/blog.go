package pojo

import (
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model `json:"-"` //gorm.Model 是一个基础结构体，包含了ID, CreatedAt, UpdatedAt, DeletedAt 等字段
	Id         int        `json:"id" `
	Title      string     `json:"title"`
	Content    string     `json:"content"`
	Category   string     `json:"category"`
}

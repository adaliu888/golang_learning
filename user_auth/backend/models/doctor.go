package models

import (
    "gorm.io/gorm"
)

type Doctor struct {
    gorm.Model
    UserID   uint   `gorm:"not null"` // 关联用户
    Name     string `gorm:"not null"`
    Specialty string `gorm:"not null"`
}
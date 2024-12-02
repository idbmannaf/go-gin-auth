package models

import (
	"time"
)

type Permission struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique"`
}

type User struct {
	ID         int    `json:"id" gorm:"primaryKey;autoIncrement"`
	UserName   string `json:"userName" gorm:"unique"`
	Password   string `json:"password" gorm:"password"`
	IsAdmin    bool   `json:"isAdmin" gorm:"default:0"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Permission []Permission `gorm:"many2many:user_permissions;"`
}

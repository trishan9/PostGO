package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"body,omitempty"`
	Password string `json:"-"`
	Avatar   string `json:"avatar,omitempty"`
	// Posts    []Post
}

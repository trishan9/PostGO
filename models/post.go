package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	UserID uint   `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body,omitempty"`
}

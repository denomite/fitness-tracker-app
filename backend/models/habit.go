package models

import "gorm.io/gorm"

type Habit struct {
	gorm.Model
	UserID      uint
	Title       string
	Description string
	IsDaily     bool
	Completed   bool
}

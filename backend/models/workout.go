package models

import "gorm.io/gorm"

type Workout struct {
	gorm.Model
	UserID   uint
	Type     string
	Duration int
	Calories int
	Notes    string
}

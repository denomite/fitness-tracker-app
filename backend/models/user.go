package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string    `json:"username" gorm:"unique"`
	Email    string    `json:"email" gorm:"unique"`
	Password string    `json:"-"`
	Habits   []Habit   `gorm:"foreignKey:UserID"`
	Meals    []Meal    `gorm:"foreignKey:UserID"`
	Workouts []Workout `gorm:"foreignKey:UserID"`
}

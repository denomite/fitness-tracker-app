package models

import "gorm.io/gorm"

type Meal struct {
	gorm.Model
	UserID   uint
	Type     string //  "breakfast", "lunch", etc.
	Calories int
	Protein  int
	Carbs    int
	Fat      int
	Notes    string
}

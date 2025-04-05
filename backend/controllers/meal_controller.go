package controllers

import (
	"fitnes-tracker/database"
	"fitnes-tracker/models"
	"fitnes-tracker/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MealInput struct {
	Type     string `json:"type" validate:"required"`
	Calories int    `json:"calories" validate:"required,min=0"`
	Protein  int    `json:"protein"`
	Carbs    int    `json:"carbs"`
	Fat      int    `json:"fat"`
	Notes    string `json:"notes"`
}

func CreateMeal(c *gin.Context) {
	var input models.Meal
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.MustGet("userID").(uint)
	input.UserID = userID

	database.GetDB().Create(&input)
	c.JSON(http.StatusCreated, input)
}

func GetMeals(c *gin.Context) {
	var meals []models.Meal
	userID := c.MustGet("userID").(uint)
	database.GetDB().Where("user_id = ?", userID).Find(&meals)
	c.JSON(http.StatusOK, meals)
}

func UpdateMeal(c *gin.Context) {
	var meal models.Meal
	id := c.Param("id")

	if err := database.GetDB().First(&meal, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Meal not found"})
		return
	}

	if meal.UserID != c.MustGet("userID").(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		return
	}

	var input MealInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.Validate.Struct(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	meal.Type = input.Type
	meal.Calories = input.Calories
	meal.Protein = input.Protein
	meal.Carbs = input.Carbs
	meal.Fat = input.Fat
	meal.Notes = input.Notes

	database.GetDB().Save(&meal)
	c.JSON(http.StatusOK, meal)
}

func DeleteMeal(c *gin.Context) {
	var meal models.Meal
	id := c.Param("id")

	if err := database.GetDB().First(&meal, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Meal not found"})
		return
	}

	if meal.UserID != c.MustGet("userID").(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		return
	}

	database.GetDB().Delete(&meal)
	c.JSON(http.StatusOK, gin.H{"message": "Meal deleted"})
}

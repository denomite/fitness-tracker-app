package controllers

import (
	"fitnes-tracker/database"
	"fitnes-tracker/models"
	"fitnes-tracker/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HabitInput struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	IsDaily     bool   `json:"is_daily"`
	Completed   bool   `json:"completed"`
}

func CreateHabit(c *gin.Context) {
	var input models.Habit
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.MustGet("userID").(uint)
	input.UserID = userID

	database.GetDB().Create(&input)
	c.JSON(http.StatusCreated, input)
}

func GetHabits(c *gin.Context) {
	var habits []models.Habit
	userID := c.MustGet("userID").(uint)
	database.GetDB().Where("user_id = ?", userID).Find(&habits)
	c.JSON(http.StatusOK, habits)
}

func UpdateHabit(c *gin.Context) {
	var habit models.Habit
	id := c.Param("id")

	if err := database.GetDB().First(&habit, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Habit not found"})
		return
	}

	if habit.UserID != c.MustGet("userID").(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		return
	}

	var input HabitInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.Validate.Struct(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	habit.Title = input.Title
	habit.Description = input.Description
	habit.IsDaily = input.IsDaily
	habit.Completed = input.Completed

	database.GetDB().Save(&habit)
	c.JSON(http.StatusOK, habit)
}

func DeleteHabit(c *gin.Context) {
	var habit models.Habit
	id := c.Param("id")

	if err := database.GetDB().First(&habit, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Habit not found"})
		return
	}

	if habit.UserID != c.MustGet("userID").(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		return
	}

	database.GetDB().Delete(&habit)
	c.JSON(http.StatusOK, gin.H{"message": "Habit deleted"})
}

package controllers

import (
	"fitnes-tracker/database"
	"fitnes-tracker/models"
	"fitnes-tracker/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WorkoutInput struct {
	Type     string `json:"type" validate:"required"`
	Duration int    `json:"duration" validate:"required,min=1"`
	Calories int    `json:"calories"`
	Notes    string `json:"notes"`
}

func CreateWorkout(c *gin.Context) {
	var input WorkoutInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.Validate.Struct(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	workout := models.Workout{
		Type:     input.Type,
		Duration: input.Duration,
		Calories: input.Calories,
		Notes:    input.Notes,
		UserID:   c.MustGet("userID").(uint),
	}

	database.GetDB().Create(&workout)
	c.JSON(http.StatusCreated, workout)
}

func GetWorkouts(c *gin.Context) {
	var workouts []models.Workout
	userID := c.MustGet("userID").(uint)
	database.GetDB().Where("user_id = ?", userID).Find(&workouts)
	c.JSON(http.StatusOK, workouts)
}

func UpdateWorkout(c *gin.Context) {
	var workout models.Workout
	id := c.Param("id")

	if err := database.GetDB().First(&workout, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Workout not found"})
		return
	}

	if workout.UserID != c.MustGet("userID").(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		return
	}

	var input WorkoutInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.Validate.Struct(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	workout.Type = input.Type
	workout.Duration = input.Duration
	workout.Calories = input.Calories
	workout.Notes = input.Notes

	database.GetDB().Save(&workout)
	c.JSON(http.StatusOK, workout)
}

func DeleteWorkout(c *gin.Context) {
	var workout models.Workout
	id := c.Param("id")

	if err := database.GetDB().First(&workout, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Workout not found"})
		return
	}

	if workout.UserID != c.MustGet("userID").(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		return
	}

	database.GetDB().Delete(&workout)
	c.JSON(http.StatusOK, gin.H{"message": "Workout deleted"})
}

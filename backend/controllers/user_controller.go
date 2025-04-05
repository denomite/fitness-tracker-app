package controllers

import (
	"fitnes-tracker/database"
	"fitnes-tracker/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetProfile - Returns the profile of the authenticated user
func GetProfile(c *gin.Context) {
	userID := c.MustGet("userID").(uint) // Get the user ID from the JWT middleware

	var user models.User
	if err := database.GetDB().Preload("Habits").Preload("Meals").Preload("Workouts").First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Return user profile (you can customize what to send back)
	c.JSON(http.StatusOK, gin.H{
		"username": user.Username,
		"email":    user.Email,
		"habits":   user.Habits,
		"meals":    user.Meals,
		"workouts": user.Workouts,
	})
}

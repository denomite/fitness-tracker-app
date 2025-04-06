package controllers

import (
	"errors"
	"fitnes-tracker/database"
	"fitnes-tracker/models"
	"fitnes-tracker/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Register handles user registration

// Register handles user registration
// func Register(c *gin.Context) {
// 	var input models.User
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data: " + err.Error()})
// 		return
// 	}

// 	// Validate input data (email, username, password)
// 	if input.Email == "" || input.Username == "" || input.Password == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Email, Username, and Password are required"})
// 		return
// 	}

// 	db := database.GetDB()
// 	var existingUser models.User

// 	// Check if email already exists
// 	if err := db.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
// 		c.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
// 		return
// 	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error (email check)"})
// 		return
// 	}

// 	// Check if username already exists
// 	if err := db.Where("username = ?", input.Username).First(&existingUser).Error; err == nil {
// 		c.JSON(http.StatusConflict, gin.H{"error": "Username already taken"})
// 		return
// 	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error (username check)"})
// 		return
// 	}

// 	// Hash the password
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
// 		return
// 	}
// 	input.Password = string(hashedPassword)

// 	// Create user
// 	if err := db.Create(&input).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
// }

// DOESNT WORK
// func Register(c *gin.Context) {
// 	var input models.User
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Hash password
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
// 		return
// 	}

// 	// Create user
// 	user := models.User{
// 		Username: input.Username,
// 		Email:    input.Email,
// 		Password: string(hashedPassword),
// 	}

// 	if err := database.GetDB().Create(&user).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, gin.H{"message": "User registered"})

// WORKS
func Register(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.GetDB()
	var existingUser models.User

	// Check if email already exists
	if err := db.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
		return
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error (email check)"})
		return
	}

	// Check if username already exists
	if err := db.Where("username = ?", input.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already taken"})
		return
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error (username check)"})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	input.Password = string(hashedPassword)

	// Create user
	if err := db.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})

}

// type LoginInput struct {
// 	Email    string `json:"email" binding:"required"`
// 	Password string `json:"password" binding:"required"`
// }

// func Login(c *gin.Context) {
// 	var input LoginInput
// 	var user models.User

// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
// 		return
// 	}

// 	fmt.Println("Login attempt for:", input.Email)

// 	db := database.GetDB()
// 	if db == nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not connected"})
// 		return
// 	}

// 	if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
// 			return
// 		}
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error: " + err.Error()})
// 		return
// 	}

// 	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
// 		return
// 	}

// 	token, err := utils.GenerateJWT(user.ID)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed: " + err.Error()})
// 		return
// 	}

// 	fmt.Println("Login successful for:", input.Email)
// 	c.JSON(http.StatusOK, gin.H{"token": token})
// }

// Login handles user login
func Login(c *gin.Context) {
	var input models.User
	var user models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Login attempt for: ", input.Email)

	// Find user by email
	if err := database.GetDB().Where("email = ?", input.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error (login)"})
		return
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate JWT
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	fmt.Println("Login successful for:", input.Email) // Debug log

	c.JSON(http.StatusOK, gin.H{"token": token})
}

package routes

import (
	"fitnes-tracker/controllers"
	"fitnes-tracker/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")

	// Auth
	api.POST("/register", controllers.Register)
	api.POST("/login", controllers.Login)

	// Protected
	auth := api.Group("/")
	auth.Use(middlewares.JWTAuthMiddleware())
	auth.GET("/profile", controllers.GetProfile)

	auth.POST("/workouts", controllers.CreateWorkout)
	auth.GET("/workouts", controllers.GetWorkouts)
	auth.PUT("/workouts/:id", controllers.UpdateWorkout)
	auth.DELETE("/workouts/:id", controllers.DeleteWorkout)

	auth.POST("/meals", controllers.CreateMeal)
	auth.GET("/meals", controllers.GetMeals)
	auth.PUT("/meals/:id", controllers.UpdateMeal)
	auth.DELETE("/meals/:id", controllers.DeleteMeal)

	auth.POST("/habits", controllers.CreateHabit)
	auth.GET("/habits", controllers.GetHabits)
	auth.PUT("/habits/:id", controllers.UpdateHabit)
	auth.DELETE("/habits/:id", controllers.DeleteHabit)
}

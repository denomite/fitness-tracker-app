package main

import (
	"fitnes-tracker/config"
	"fitnes-tracker/database"
	"fitnes-tracker/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	config.Loadenv()

	database.ConnectDatabase()

	r := gin.Default()

	routes.RegisterRoutes(r)

	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Error starting the server: ", err)
	}
	fmt.Println("Server started on http://127.0.0.1:8080")
}

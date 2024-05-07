package main

import (
	"fmt"

	"github.com/akashdesale98/GolangAssessment/handler"
	service "github.com/akashdesale98/GolangAssessment/service"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Initialize database
	db, err := service.NewPostgresDB()
	if err != nil {
		panic(err)
	}
	fmt.Println("database connection initiated")

	conn := &handler.Handler{
		DB: db,
	}

	// Define routes
	router.GET("/employee", conn.GetEmployeeByID)

	// Start the server
	router.Run(":8080")
}

package main

import (
	"fmt"

	"github.com/akashdesale98/GolangAssessment/handler"
	service "github.com/akashdesale98/GolangAssessment/service"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Initialize db coonnection
	db, err := service.NewPostgresDB()
	if err != nil {
		panic(err)
	}
	fmt.Println("database connection initiated")

	conn := &handler.Handler{
		DB: db,
	}

	//routes
	router.POST("/employee", conn.CreateEmployee)
	router.GET("/employee", conn.GetEmployeeByID)
	router.PUT("/employee", conn.UpdateEmployee)
	router.GET("/employees", conn.GetEmployees)
	router.DELETE("/employee/:id", conn.DeleteEmployee)

	// Start the server
	router.Run(":8080")
}

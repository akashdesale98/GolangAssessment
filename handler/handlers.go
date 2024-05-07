package handler

import (
	"log"
	"net/http"

	"github.com/akashdesale98/GolangAssessment/service"
	"github.com/gin-gonic/gin"
)

// PostgreSQL implementation of the Database interface
type Handler struct {
	DB *service.PostgresDB
}

func (h *Handler) CreateEmployee(c *gin.Context) {
	// Implementation to add a new employee to the database

	c.JSON(http.StatusOK, "Hello World")
}

func (h *Handler) GetEmployeeByID(c *gin.Context) {
	employeeID := c.Query("id")
	emp, err := h.DB.GetEmployeeByID(c.Request.Context(), employeeID)
	if err != nil {
		log.Println("emp", employeeID, "err", err)
		c.AbortWithStatus(http.StatusNotImplemented)
	}

	c.JSON(http.StatusOK, emp)
}

func (h *Handler) UpdateEmployee(c *gin.Context) {
	// Implementation to update an existing employee in the database
	c.JSON(http.StatusOK, "Hello World")
}

func (h *Handler) DeleteEmployee(c *gin.Context) {
	// Implementation to delete an employee by ID from the database
	c.JSON(http.StatusOK, "Hello World")
}

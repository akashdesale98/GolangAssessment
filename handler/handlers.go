package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/akashdesale98/GolangAssessment/constant"
	"github.com/akashdesale98/GolangAssessment/model"

	"github.com/akashdesale98/GolangAssessment/service"
	"github.com/gin-gonic/gin"
)

// PostgreSQL implementation of the Database interface
type Handler struct {
	DB *service.PostgresDB
}

func (h *Handler) CreateEmployee(c *gin.Context) {

	// Parse request body to extract employee details to be inserted
	var emp model.Employee
	if err := c.BindJSON(&emp); err != nil {
		log.Println("emp", emp, "err", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.DB.CreateEmployee(c.Request.Context(), &emp)
	if err != nil {
		log.Println("emp", emp, "err", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Employee inserted successfully",
	})
}

func (h *Handler) GetEmployeeByID(c *gin.Context) {
	employeeID := c.Query("id")
	emp, err := h.DB.GetEmployeeByID(c.Request.Context(), employeeID)
	if err != nil {
		log.Println("employeeID", employeeID, "err", err)
		if err.Error() == constant.ErrNoRecordPresent.Error() {
			c.JSON(http.StatusNotFound, gin.H{"error": "Employee Not Found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, emp)
}

func (h *Handler) UpdateEmployee(c *gin.Context) {

	// Parse request body to extract employee details to be updated
	var emp model.Employee
	if err := c.BindJSON(&emp); err != nil {
		log.Println("employeeID", emp.ID, "err", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check if employee to be updated exists or not
	_, err := h.DB.GetEmployeeByID(c.Request.Context(), emp.ID)
	if err != nil {
		log.Println("employeeID", emp.ID, "err", err)
		if err.Error() == constant.ErrNoRecordPresent.Error() {
			c.JSON(http.StatusNotFound, gin.H{"error": "Employee Not Found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.DB.UpdateEmployee(c.Request.Context(), &emp)
	if err != nil {
		log.Println("employeeID", emp.ID, "err", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee details updated successfully"})
}

func (h *Handler) DeleteEmployee(c *gin.Context) {
	// Get employee ID from URL parameter
	employeeID := c.Param("id")

	// check if employee to be deleted exists or not
	_, err := h.DB.GetEmployeeByID(c.Request.Context(), employeeID)
	if err != nil {
		log.Println("employeeID", employeeID, "err", err)
		if err.Error() == constant.ErrNoRecordPresent.Error() {
			c.JSON(http.StatusNotFound, gin.H{"error": "Employee Not Found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.DB.DeleteEmployee(c.Request.Context(), employeeID)
	if err != nil {
		log.Println("employeeID", employeeID, "err", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Employee with ID %s deleted successfully", employeeID),
	})
}

func (h *Handler) GetEmployees(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 1 // Default to page 1 if not provided
	}
	limit, _ := strconv.Atoi(c.Query("limit"))
	if limit == 0 {
		limit = 10 // Default limit to 10 if not provided
	}

	offset := (page - 1) * limit

	emps, err := h.DB.GetEmployees(c.Request.Context(), limit, offset)
	if err != nil {
		log.Println("GetEmployees Err :: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, emps)
}

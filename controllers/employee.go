package controllers

import (
	"EmployeeCRUD/db"
	"EmployeeCRUD/forms"
	"EmployeeCRUD/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type EmployeeController struct{}

func (u EmployeeController) Retrieve(c *gin.Context) {
	if c.Param("id") != "" {
		employee := db.Retrieve(c.Param("id"))
		if employee == (models.Employee{}) {
			c.JSON(http.StatusNotFound, gin.H{"message": "Employee not found!"})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Employee founded!", "employee": employee})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
	return
}

func (u EmployeeController) Create(c *gin.Context) {
	var employeeSignUp forms.EmployeeSignUp
	if err := c.ShouldBindJSON(&employeeSignUp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request", "error": err})
		c.Abort()
		return
	}
	employee := models.Employee{
		ID:       uuid.New().String(),
		Name:     employeeSignUp.Name,
		Position: employeeSignUp.Position,
		Salary:   employeeSignUp.Salary,
	}
	db.Save(employee.ID, employee)
	c.JSON(http.StatusCreated, gin.H{"message": "Employee created!", "employee": employee})
	return
}

func (u EmployeeController) Delete(c *gin.Context) {
	if c.Param("id") != "" {
		db.Delete(c.Param("id"))
		c.JSON(http.StatusOK, gin.H{"message": "Employee deleted!"})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
	return
}

func (u EmployeeController) Update(c *gin.Context) {
	var employeeUpdate forms.EmployeeUpdate
	if c.Param("id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request path", "error": "id is required"})
		c.Abort()
		return
	}
	if err := c.ShouldBindJSON(&employeeUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body", "error": err})
		c.Abort()
		return
	}
	employee := models.Employee{
		ID:       c.Param("id"),
		Name:     employeeUpdate.Name,
		Position: employeeUpdate.Position,
		Salary:   employeeUpdate.Salary,
	}
	if db.Retrieve(employee.ID) == (models.Employee{}) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Employee not found!", "error": "insert a valid id"})
		c.Abort()
		return
	}
	db.Update(employee.ID, employee)
	c.JSON(http.StatusOK, gin.H{"message": "Employee updated!", "employee": employee})
	return
}

func (u EmployeeController) List(c *gin.Context) {
	pagination := forms.Pagination{}
	if err := c.ShouldBindQuery(&pagination); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request", "error": err})
		c.Abort()
		return
	}
	employees := db.List(pagination.Page, pagination.Limit)
	c.JSON(http.StatusOK, gin.H{"message": "Employees founded!", "employees": employees})
	return
}

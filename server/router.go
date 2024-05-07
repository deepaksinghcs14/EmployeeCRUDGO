package server

import (
	"EmployeeCRUD/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)

	v1 := router.Group("v1")
	{
		employeeGroup := v1.Group("employee")
		{
			employee := new(controllers.EmployeeController)
			employeeGroup.GET("/:id", employee.Retrieve)
			employeeGroup.POST("/", employee.Create)
			employeeGroup.DELETE("/:id", employee.Delete)
			employeeGroup.PUT("/:id", employee.Update)
			employeeGroup.GET("", employee.List)
		}
	}
	return router

}

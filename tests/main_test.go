package tests

import (
	"EmployeeCRUD/config"
	"EmployeeCRUD/controllers"
	"EmployeeCRUD/models"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type EmployeeSuite struct {
	config *viper.Viper
	router *gin.Engine
}

type Response struct {
	Message  string          `json:"message"`
	Employee models.Employee `json:"employee"`
}

type ListResponse struct {
	Message   string            `json:"message"`
	Employees []models.Employee `json:"employees"`
}

func (s *EmployeeSuite) SetUpTest() {
	config.Init("test")
	s.config = config.GetConfig()
	s.router = SetupRouter()
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	gin.SetMode(gin.TestMode)
	health := new(controllers.HealthController)
	router.GET("/health", health.Status)
	v1 := router.Group("v1")
	{
		employeeGroup := v1.Group("employee")
		{
			employee := new(controllers.EmployeeController)
			employeeGroup.GET("/:id", employee.Retrieve)
			employeeGroup.POST("", employee.Create)
			employeeGroup.DELETE("/:id", employee.Delete)
			employeeGroup.PUT("/:id", employee.Update)
			employeeGroup.GET("", employee.List)
		}
	}
	return router
}

func TestEmployeeCrud(t *testing.T) {
	var jsonStr = []byte(`{"name":"test","position":"test","salary":1000}`)

	// create employee
	req, _ := http.NewRequest("POST", "/v1/employee", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router := SetupRouter()
	router.ServeHTTP(resp, req)
	assert.Equal(t, 201, resp.Code)
	body := resp.Body.String()
	assert.NotEmpty(t, body)
	// get employee from body
	response := Response{}
	err := json.Unmarshal([]byte(body), &response)
	employee := response.Employee
	assert.Equal(t, "test", employee.Name)

	// get employee
	id := employee.ID
	req, _ = http.NewRequest("GET", "/v1/employee/"+id, nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code)
	body = resp.Body.String()
	assert.NotEmpty(t, body)
	response = Response{}
	err = json.Unmarshal([]byte(body), &response)
	employee = response.Employee
	assert.Nil(t, err)
	assert.Equal(t, "test", employee.Name)
	assert.Equal(t, id, employee.ID)

	// update employee
	jsonStr = []byte(`{"name":"test2","position":"test2","salary":2000}`)
	req, _ = http.NewRequest("PUT", "/v1/employee/"+id, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code)
	body = resp.Body.String()
	assert.NotEmpty(t, body)
	response = Response{}
	err = json.Unmarshal([]byte(body), &response)
	employee = response.Employee
	assert.Equal(t, "test2", employee.Name)
	assert.Equal(t, float64(2000), employee.Salary)
	assert.Equal(t, id, employee.ID)

	// delete employee
	req, _ = http.NewRequest("DELETE", "/v1/employee/"+id, nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code)
	body = resp.Body.String()
	response = Response{}
	err = json.Unmarshal([]byte(body), &response)
	assert.Equal(t, "Employee deleted!", response.Message)

	// get employee
	req, _ = http.NewRequest("GET", "/v1/employee/"+id, nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, 404, resp.Code)
	body = resp.Body.String()
	assert.NotEmpty(t, body)
	response = Response{}
	err = json.Unmarshal([]byte(body), &response)
	assert.Equal(t, "Employee not found!", response.Message)
}

func TestListEmployee(t *testing.T) {
	// create 100 employees
	router := SetupRouter()
	for i := 0; i < 100; i++ {
		jsonStr := []byte(`{"name":"test","position":"test","salary":1000}`)
		req, _ := http.NewRequest("POST", "/v1/employee", bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		assert.Equal(t, 201, resp.Code)
	}

	// list employees
	req, _ := http.NewRequest("GET", "/v1/employee?page=1&limit=10", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code)
	body := resp.Body.String()
	response := ListResponse{}
	err := json.Unmarshal([]byte(body), &response)
	assert.Nil(t, err)
	assert.Equal(t, 10, len(response.Employees))

}

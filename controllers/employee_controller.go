package controller

import (
	"juke-test-restapi/exception"
	"juke-test-restapi/model"
	"juke-test-restapi/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetAllEmployees(c *gin.Context) {
	zap.L().Info("GET /api/employees")
	employees, err := service.GetAllEmployees(c.Request.Context())
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"data": employees})
}

func GetEmployeeByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		panic(exception.ErrBadRequest)
	}
	zap.L().Info("GET /api/employees/:id", zap.Int64("id", id))

	employee, err := service.GetEmployeeByID(c.Request.Context(), id)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"data": employee})
}

func CreateEmployee(c *gin.Context) {
	var req model.CreateEmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(exception.NewAppError(400, err.Error()))
	}
	zap.L().Info("POST /api/employees", zap.String("email", req.Email))

	employee, err := service.CreateEmployee(c.Request.Context(), req)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusCreated, gin.H{"data": employee})
}

func UpdateEmployee(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		panic(exception.ErrBadRequest)
	}

	var req model.UpdateEmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(exception.NewAppError(400, err.Error()))
	}
	zap.L().Info("PUT /api/employees/:id", zap.Int64("id", id), zap.String("email", req.Email))

	employee, err := service.UpdateEmployee(c.Request.Context(), id, req)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"data": employee})
}

func DeleteEmployee(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		panic(exception.ErrBadRequest)
	}
	zap.L().Info("DELETE /api/employees/:id", zap.Int64("id", id))

	if err := service.DeleteEmployee(c.Request.Context(), id); err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}

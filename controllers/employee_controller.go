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

// GetAllEmployees godoc
// @Summary Get all employees
// @Description Get list of all employees
// @Tags employees
// @Produce json
// @Success 200 {object} map[string][]model.Employee
// @Router /employees [get]
func GetAllEmployees(c *gin.Context) {
	zap.L().Info("GET /api/employees")
	employees, err := service.GetAllEmployees(c.Request.Context())
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"data": employees})
}

// GetEmployeeByID godoc
// @Summary Get employee by ID
// @Description Get employee details by ID
// @Tags employees
// @Produce json
// @Param id path int true "Employee ID"
// @Success 200 {object} map[string]model.Employee
// @Failure 404 {object} map[string]interface{}
// @Router /employees/{id} [get]
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

// CreateEmployee godoc
// @Summary Create new employee
// @Description Create a new employee record
// @Tags employees
// @Accept json
// @Produce json
// @Param employee body model.CreateEmployeeRequest true "Employee data"
// @Success 201 {object} map[string]model.Employee
// @Failure 400 {object} map[string]interface{}
// @Router /employees [post]
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

// UpdateEmployee godoc
// @Summary Update employee
// @Description Update employee data by ID
// @Tags employees
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Param employee body model.UpdateEmployeeRequest true "Employee data"
// @Success 200 {object} map[string]model.Employee
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /employees/{id} [put]
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

// DeleteEmployee godoc
// @Summary Delete employee
// @Description Delete employee by ID
// @Tags employees
// @Produce json
// @Param id path int true "Employee ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]interface{}
// @Router /employees/{id} [delete]
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

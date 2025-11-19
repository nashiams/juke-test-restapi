package router

import (
	controller "juke-test-restapi/controllers"
	"juke-test-restapi/exception"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    r.Use(exception.ErrorHandler())

    r.GET("/api/employees", controller.GetAllEmployees)
    r.GET("/api/employees/:id", controller.GetEmployeeByID)
    r.POST("/api/employees", controller.CreateEmployee)
    r.PUT("/api/employees/:id", controller.UpdateEmployee)
    r.DELETE("/api/employees/:id", controller.DeleteEmployee)

    return r
}

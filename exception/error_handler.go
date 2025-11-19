package exception

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *AppError) Error() string {
	return e.Message
}

// Predefined errors
var (
	ErrNotFound     = &AppError{Code: http.StatusNotFound, Message: "Resource not found"}
	ErrBadRequest   = &AppError{Code: http.StatusBadRequest, Message: "Bad request"}
	ErrUnauthorized = &AppError{Code: http.StatusUnauthorized, Message: "Unauthorized"}
	ErrInternal     = &AppError{Code: http.StatusInternalServerError, Message: "Internal server error"}
)

func NewAppError(code int, message string) *AppError {
	return &AppError{Code: code, Message: message}
}

// ErrorHandler middleware
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch e := err.(type) {
				case *AppError:
					c.JSON(e.Code, gin.H{
						"error":   true,
						"message": e.Message,
						"code":    e.Code,
					})
				case error:
					c.JSON(http.StatusInternalServerError, gin.H{
						"error":   true,
						"message": e.Error(),
						"code":    http.StatusInternalServerError,
					})
				default:
					c.JSON(http.StatusInternalServerError, gin.H{
						"error":   true,
						"message": "Internal server error",
						"code":    http.StatusInternalServerError,
					})
				}
				c.Abort()
			}
		}()
		c.Next()

		// Handle errors set in context
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			if appErr, ok := err.(*AppError); ok {
				c.JSON(appErr.Code, gin.H{
					"error":   true,
					"message": appErr.Message,
					"code":    appErr.Code,
				})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   true,
					"message": err.Error(),
					"code":    http.StatusInternalServerError,
				})
			}
		}
	}
}

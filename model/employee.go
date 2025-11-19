package model

import "time"

type Employee struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Position  string    `json:"position"`
	Salary    float64   `json:"salary"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateEmployeeRequest struct {
	Name     string  `json:"name" binding:"required"`
	Email    string  `json:"email" binding:"required,email"`
	Position string  `json:"position" binding:"required"`
	Salary   float64 `json:"salary" binding:"required,gt=0"`
}

type UpdateEmployeeRequest struct {
	Name     string  `json:"name" binding:"required"`
	Email    string  `json:"email" binding:"required,email"`
	Position string  `json:"position" binding:"required"`
	Salary   float64 `json:"salary" binding:"required,gt=0"`
}

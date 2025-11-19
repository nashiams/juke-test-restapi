package repository

import (
	"context"
	"juke-test-restapi/db"
	"juke-test-restapi/model"

	"go.uber.org/zap"
)

func GetAll(ctx context.Context) ([]model.Employee, error) {
	query := `SELECT id, name, email, position, salary, created_at FROM employees ORDER BY id`
	zap.L().Debug("SQL: SELECT all employees")
	rows, err := db.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []model.Employee
	for rows.Next() {
		var emp model.Employee
		if err := rows.Scan(&emp.ID, &emp.Name, &emp.Email, &emp.Position, &emp.Salary, &emp.CreatedAt); err != nil {
			return nil, err
		}
		employees = append(employees, emp)
	}
	return employees, nil
}

func GetByID(ctx context.Context, id int64) (*model.Employee, error) {
	query := `SELECT id, name, email, position, salary, created_at FROM employees WHERE id = $1`
	zap.L().Debug("SQL: SELECT employee by ID", zap.Int64("id", id))
	var emp model.Employee
	err := db.DB.QueryRow(ctx, query, id).Scan(&emp.ID, &emp.Name, &emp.Email, &emp.Position, &emp.Salary, &emp.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &emp, nil
}

func Create(ctx context.Context, req model.CreateEmployeeRequest) (*model.Employee, error) {
	query := `INSERT INTO employees (name, email, position, salary) VALUES ($1, $2, $3, $4) RETURNING id, name, email, position, salary, created_at`
	zap.L().Debug("SQL: INSERT employee", zap.String("email", req.Email))
	var emp model.Employee
	err := db.DB.QueryRow(ctx, query, req.Name, req.Email, req.Position, req.Salary).Scan(&emp.ID, &emp.Name, &emp.Email, &emp.Position, &emp.Salary, &emp.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &emp, nil
}

func Update(ctx context.Context, id int64, req model.UpdateEmployeeRequest) (*model.Employee, error) {
	query := `UPDATE employees SET name = $1, email = $2, position = $3, salary = $4 WHERE id = $5 RETURNING id, name, email, position, salary, created_at`
	zap.L().Debug("SQL: UPDATE employee", zap.Int64("id", id))
	var emp model.Employee
	err := db.DB.QueryRow(ctx, query, req.Name, req.Email, req.Position, req.Salary, id).Scan(&emp.ID, &emp.Name, &emp.Email, &emp.Position, &emp.Salary, &emp.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &emp, nil
}

func Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM employees WHERE id = $1`
	zap.L().Debug("SQL: DELETE employee", zap.Int64("id", id))
	result, err := db.DB.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return nil
	}
	return nil
}

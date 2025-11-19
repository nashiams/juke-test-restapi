package service

import (
	"context"
	"juke-test-restapi/exception"
	"juke-test-restapi/model"
	"juke-test-restapi/repository"
	"strings"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

func GetAllEmployees(ctx context.Context) ([]model.Employee, error) {
	return repository.GetAll(ctx)
}

func GetEmployeeByID(ctx context.Context, id int64) (*model.Employee, error) {
	emp, err := repository.GetByID(ctx, id)
	if err == pgx.ErrNoRows {
		zap.L().Warn("Employee not found", zap.Int64("id", id))
		return nil, exception.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return emp, nil
}

func CreateEmployee(ctx context.Context, req model.CreateEmployeeRequest) (*model.Employee, error) {
	emp, err := repository.Create(ctx, req)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "unique") {
			zap.L().Warn("Duplicate email", zap.String("email", req.Email))
			return nil, exception.NewAppError(400, "Email already exists")
		}
		return nil, err
	}
	return emp, nil
}

func UpdateEmployee(ctx context.Context, id int64, req model.UpdateEmployeeRequest) (*model.Employee, error) {
	emp, err := repository.Update(ctx, id, req)
	if err == pgx.ErrNoRows {
		zap.L().Warn("Employee not found", zap.Int64("id", id))
		return nil, exception.ErrNotFound
	}
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "unique") {
			zap.L().Warn("Duplicate email", zap.String("email", req.Email))
			return nil, exception.NewAppError(400, "Email already exists")
		}
		return nil, err
	}
	return emp, nil
}

func DeleteEmployee(ctx context.Context, id int64) error {
	_, err := repository.GetByID(ctx, id)
	if err == pgx.ErrNoRows {
		zap.L().Warn("Employee not found", zap.Int64("id", id))
		return exception.ErrNotFound
	}
	if err != nil {
		return err
	}
	return repository.Delete(ctx, id)
}

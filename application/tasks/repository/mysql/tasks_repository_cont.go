package repository

import "Bank-INA/domain"

type TasksRepositoryContract interface {
	CreateTasks(data domain.Tasks) (err domain.ErrorData)
	GetAll() (result []domain.Tasks, err domain.ErrorData)
	GetByID(id int) (result domain.Tasks, err domain.ErrorData)
	UpdateByID(id int, updates map[string]interface{}) (err domain.ErrorData)
	DeleteByID(id int) (err domain.ErrorData)
}

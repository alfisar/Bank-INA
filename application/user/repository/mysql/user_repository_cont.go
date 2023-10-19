package repository

import "Bank-INA/domain"

type UserRepositoryContract interface {
	CreateUser(data domain.User) (err domain.ErrorData)
	GetAll() (result []domain.User, err domain.ErrorData)
	GetByID(id int) (result domain.User, err domain.ErrorData)
	UpdateByID(id int, updates map[string]interface{}) (err domain.ErrorData)
	DeleteByID(id int) (err domain.ErrorData)
}

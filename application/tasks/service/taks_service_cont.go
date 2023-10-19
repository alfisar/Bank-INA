package service

import "Bank-INA/domain"

type TaskserviceContract interface {
	Create(data domain.Tasks) (err domain.ErrorData)
	GetAll() (result []domain.Tasks, err domain.ErrorData)
	GetOne(ID int) (result domain.Tasks, err domain.ErrorData)
	Update(data domain.Tasks) (err domain.ErrorData)
	Delete(ID int) (err domain.ErrorData)
}

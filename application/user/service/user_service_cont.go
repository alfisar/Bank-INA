package service

import "Bank-INA/domain"

type UserviceContract interface {
	Create(data domain.User) (err domain.ErrorData)
	GetAll() (result []domain.User, err domain.ErrorData)
	GetOne(ID int) (result domain.User, err domain.ErrorData)
	Update(data domain.User) (err domain.ErrorData)
	Delete(ID int) (err domain.ErrorData)
}

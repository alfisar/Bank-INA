package service

import (
	repo "Bank-INA/application/user/repository/mysql"
	"Bank-INA/domain"
	"Bank-INA/internal/errorhandler"
	"Bank-INA/internal/hashing"
	"fmt"
)

type UserService struct {
	repo repo.UserRepositoryContract
}

func NewUserService(repo repo.UserRepositoryContract) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (obj *UserService) Create(data domain.User) (err domain.ErrorData) {
	data.Password, err = hashing.GeneratePass(data.Password)
	if err.ResponseCode != "" {
		return
	}
	return obj.repo.CreateUser(data)
}

func (obj *UserService) GetAll() (result []domain.User, err domain.ErrorData) {
	return obj.repo.GetAll()
}

func (obj *UserService) GetOne(ID int) (result domain.User, err domain.ErrorData) {
	return obj.repo.GetByID(ID)
}

func (obj *UserService) Update(data domain.User) (err domain.ErrorData) {
	updates := map[string]interface{}{}

	if data.Name != "" {
		updates["name"] = data.Name
	}

	if data.Email != "" {
		updates["email"] = data.Email
	}

	if updates == nil {
		err = errorhandler.FailedFromService(domain.FailedServ, fmt.Errorf("Data update is empty"))
		return
	}
	err = obj.repo.UpdateByID(data.ID, updates)
	return
}

func (obj *UserService) Delete(ID int) (err domain.ErrorData) {
	return obj.repo.DeleteByID(ID)
}

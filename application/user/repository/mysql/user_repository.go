package repository

import (
	"Bank-INA/domain"
	"Bank-INA/internal/errorhandler"

	"gorm.io/gorm"
)

type UserRepository struct {
	conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) *UserRepository {
	return &UserRepository{
		conn: conn.Debug(),
	}
}

func (obj *UserRepository) CreateUser(data domain.User) (err domain.ErrorData) {
	errData := obj.conn.Table("users").Create(&data).Error

	if errData != nil {
		err = errorhandler.FailedFromRepo(errData)
	}

	return
}

func (obj *UserRepository) GetAll() (result []domain.User, err domain.ErrorData) {
	errData := obj.conn.Table("users").Find(&result).Error

	if errData != nil {
		err = errorhandler.FailedFromRepo(errData)
	}

	return
}

func (obj *UserRepository) GetByID(id int) (result domain.User, err domain.ErrorData) {
	errData := obj.conn.Table("users").Where("id = ?", id).First(&result).Error

	if errData != nil {
		err = errorhandler.FailedFromRepo(errData)
	}

	return
}

func (obj *UserRepository) UpdateByID(id int, updates map[string]interface{}) (err domain.ErrorData) {
	errData := obj.conn.Table("users").Where("id = ?", id).Updates(&updates).Error
	if errData != nil {
		err = errorhandler.FailedFromRepo(errData)
	}
	return
}

func (obj *UserRepository) DeleteByID(id int) (err domain.ErrorData) {
	errData := obj.conn.Table("users").Where("id = ?", id).Delete(&domain.User{}).Error

	if errData != nil {
		err = errorhandler.FailedFromRepo(errData)
	}

	return
}

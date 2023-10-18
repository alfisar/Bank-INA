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
	errData := obj.conn.Table("user").Create(&data).Error

	if errData != nil {
		err = errorhandler.FailedFromRepo(errData)
	}

	return
}

func (obj *UserRepository) GetAll() (result []domain.User, err domain.ErrorData) {
	errData := obj.conn.Table("user").Find(&result).Error

	if errData != nil {
		err = errorhandler.FailedFromRepo(errData)
	}

	return
}

func (obj *UserRepository) GetById(id int) (result domain.User, err domain.ErrorData) {
	errData := obj.conn.Table("user").Where("id = ?", id).First(&result).Error

	if errData != nil {
		err = errorhandler.FailedFromRepo(errData)
	}

	return
}

func (obj *UserRepository) DeleteByID(id int) (err domain.ErrorData) {
	errData := obj.conn.Table("user").Where("id = ?", id).Delete(&domain.User{}).Error

	if errData != nil {
		err = errorhandler.FailedFromRepo(errData)
	}

	return
}

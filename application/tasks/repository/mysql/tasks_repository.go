package repository

import (
	"Bank-INA/domain"
	"Bank-INA/internal/errorhandler"

	"gorm.io/gorm"
)

type TaskRepository struct {
	conn *gorm.DB
}

func NewTaskRepository(conn *gorm.DB) *TaskRepository {
	return &TaskRepository{
		conn: conn.Debug(),
	}
}

func (obj *TaskRepository) CreateTasks(data domain.Tasks) (err domain.ErrorData) {
	errData := obj.conn.Table("tasks").Create(&data).Error

	if errData != nil {
		err = errorhandler.FailedFromRepo(errData)
	}

	return
}

func (obj *TaskRepository) GetAll() (result []domain.Tasks, err domain.ErrorData) {
	errData := obj.conn.Table("tasks").Find(&result).Error

	if errData != nil {
		err = errorhandler.FailedFromRepo(errData)
	}

	return
}

func (obj *TaskRepository) GetByID(id int) (result domain.Tasks, err domain.ErrorData) {
	errData := obj.conn.Table("tasks").Where("id = ?", id).First(&result).Error

	if errData != nil {
		err = errorhandler.FailedFirstFromRepo(errData)
	}

	return
}

func (obj *TaskRepository) UpdateByID(id int, updates map[string]interface{}) (err domain.ErrorData) {
	errData := obj.conn.Table("tasks").Where("id = ?", id).Updates(&updates).Error
	if errData != nil {
		err = errorhandler.FailedFromRepo(errData)
	}
	return
}

func (obj *TaskRepository) DeleteByID(id int) (err domain.ErrorData) {
	errData := obj.conn.Table("tasks").Where("id = ?", id).Delete(&domain.User{}).Error

	if errData != nil {
		err = errorhandler.FailedFromRepo(errData)
	}

	return
}

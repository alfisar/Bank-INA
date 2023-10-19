package service

import (
	repo "Bank-INA/application/tasks/repository/mysql"
	"Bank-INA/domain"
	"Bank-INA/internal/errorhandler"
	"fmt"
)

type TasksService struct {
	repo repo.TasksRepositoryContract
}

func NewTasksService(repo repo.TasksRepositoryContract) *TasksService {
	return &TasksService{
		repo: repo,
	}
}

func (obj *TasksService) Create(data domain.Tasks) (err domain.ErrorData) {
	return obj.repo.CreateTasks(data)
}

func (obj *TasksService) GetAll() (result []domain.Tasks, err domain.ErrorData) {
	return obj.repo.GetAll()
}

func (obj *TasksService) GetOne(ID int) (result domain.Tasks, err domain.ErrorData) {
	return obj.repo.GetByID(ID)
}

func (obj *TasksService) Update(data domain.Tasks) (err domain.ErrorData) {
	updates := map[string]interface{}{}

	if data.Title != "" {
		updates["title"] = data.Title
	}

	if data.Desc != "" {
		updates["description"] = data.Desc
	}

	if data.Status != "" {
		updates["status"] = data.Status
	}

	if updates == nil {
		err = errorhandler.FailedFromService(domain.FailedServ, fmt.Errorf("Data update is empty"))
		return
	}
	err = obj.repo.UpdateByID(data.ID, updates)
	return
}

func (obj *TasksService) Delete(ID int) (err domain.ErrorData) {
	return obj.repo.DeleteByID(ID)
}

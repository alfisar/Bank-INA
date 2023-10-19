package controller

import (
	"Bank-INA/application/tasks/service"
	"Bank-INA/domain"
	"Bank-INA/internal/errorhandler"
	"Bank-INA/internal/response"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TasksController struct {
	serv service.TaskserviceContract
}

func NewTasksController(serv service.TaskserviceContract) *TasksController {
	return &TasksController{
		serv: serv,
	}
}

func (obj TasksController) Create(ctx *gin.Context) {
	data := domain.Tasks{}
	errData := ctx.Bind(&data)
	if errData != nil {
		fmt.Errorf("Invalid Parse data")
		err := errorhandler.Failed(domain.FailedControllCode, domain.FailedControll, errData)
		response.Failed(ctx, err)
		return
	}

	err := obj.serv.Create(data)
	if err.ResponseCode != "" {
		response.Failed(ctx, err)
		return
	}

	response.Success(ctx, nil)
}

func (obj TasksController) Get(ctx *gin.Context) {
	data, err := obj.serv.GetAll()
	if err.ResponseCode != "" {
		response.Failed(ctx, err)
		return
	}

	response.Success(ctx, data)
}

func (obj TasksController) GetByID(ctx *gin.Context) {
	IDString := ctx.Param("id")
	ID, errData := strconv.Atoi(IDString)
	if errData != nil {
		fmt.Errorf("Invalid Parse data")
		err := errorhandler.Failed(domain.FailedControllCode, domain.FailedControll, errData)
		response.Failed(ctx, err)
		return
	}

	data, err := obj.serv.GetOne(ID)
	if err.ResponseCode != "" {
		response.Failed(ctx, err)
		return
	}

	response.Success(ctx, data)
}

func (obj TasksController) Update(ctx *gin.Context) {
	IDString := ctx.Param("id")
	ID, errData := strconv.Atoi(IDString)
	if errData != nil {
		fmt.Errorf("Invalid Parse ID")
		err := errorhandler.Failed(domain.FailedControllCode, domain.FailedControll, errData)
		response.Failed(ctx, err)
		return
	}

	data := domain.Tasks{}
	errData = ctx.Bind(&data)
	if errData != nil {
		fmt.Errorf("Invalid Parse data")
		err := errorhandler.Failed(domain.FailedControllCode, domain.FailedControll, errData)
		response.Failed(ctx, err)
		return
	}

	data.ID = ID
	err := obj.serv.Update(data)
	if err.ResponseCode != "" {
		response.Failed(ctx, err)
		return
	}

	response.Success(ctx, data)
}

func (obj TasksController) Delete(ctx *gin.Context) {
	IDString := ctx.Param("id")
	ID, errData := strconv.Atoi(IDString)
	if errData != nil {
		fmt.Errorf("Invalid Parse data")
		err := errorhandler.Failed(domain.FailedControllCode, domain.FailedControll, errData)
		response.Failed(ctx, err)
		return
	}

	err := obj.serv.Delete(ID)
	if err.ResponseCode != "" {
		response.Failed(ctx, err)
		return
	}

	response.Success(ctx, nil)
}

package controller

import (
	"Bank-INA/application/user/service"
	"Bank-INA/domain"
	"Bank-INA/internal/errorhandler"
	"Bank-INA/internal/helper"
	"Bank-INA/internal/response"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	serv service.UserviceContract
}

func NewUserController(serv service.UserviceContract) *UserController {
	return &UserController{
		serv: serv,
	}
}

func (obj UserController) Create(ctx *gin.Context) {
	data := domain.User{}
	errData := ctx.Bind(&data)
	if errData != nil {
		fmt.Errorf("Invalid Parse data")
		err := errorhandler.Failed(domain.FailedControllCode, domain.FailedControll, errData)
		response.Failed(ctx, err)
		return
	}

	err := helper.ValidateUser(data)
	if err.ResponseCode != "" {
		response.Failed(ctx, err)
		return
	}

	err = obj.serv.Create(data)
	if err.ResponseCode != "" {
		response.Failed(ctx, err)
		return
	}

	response.Success(ctx, nil)
}

func (obj UserController) Get(ctx *gin.Context) {
	data, err := obj.serv.GetAll()
	if err.ResponseCode != "" {
		response.Failed(ctx, err)
		return
	}

	response.Success(ctx, data)
}

func (obj UserController) GetByID(ctx *gin.Context) {
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

func (obj UserController) Update(ctx *gin.Context) {
	IDString := ctx.Param("id")
	ID, errData := strconv.Atoi(IDString)
	if errData != nil {
		fmt.Errorf("Invalid Parse ID")
		err := errorhandler.Failed(domain.FailedControllCode, domain.FailedControll, errData)
		response.Failed(ctx, err)
		return
	}

	data := domain.User{}
	errData = ctx.Bind(&data)
	if errData != nil {
		fmt.Errorf("Invalid Parse data")
		err := errorhandler.Failed(domain.FailedControllCode, domain.FailedControll, errData)
		response.Failed(ctx, err)
		return
	}

	if data.Email != "" {
		err := helper.ValidateUser(data)
		if err.ResponseCode != "" {
			response.Failed(ctx, err)
			return
		}
	}

	data.ID = ID
	err := obj.serv.Update(data)
	if err.ResponseCode != "" {
		response.Failed(ctx, err)
		return
	}

	response.Success(ctx, data)
}

func (obj UserController) Delete(ctx *gin.Context) {
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

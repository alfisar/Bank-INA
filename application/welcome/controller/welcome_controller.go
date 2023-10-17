package controller

import (
	"Bank-INA/internal/response"

	"github.com/gin-gonic/gin"
)

type ControllerWelcome struct {
}

func NewWelcomeController() *ControllerWelcome {
	return &ControllerWelcome{}
}

func (obj *ControllerWelcome) Hello(ctx *gin.Context) {
	response.Success(ctx, "Welcome to my API, this API for test recruiment bank INA with role Golang Developer, enjoy and cheers :) ....  ")
}

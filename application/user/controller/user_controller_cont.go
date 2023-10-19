package controller

import "github.com/gin-gonic/gin"

type UserControllerContract interface {
	Create(ctx *gin.Context)
	Get(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

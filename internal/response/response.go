package response

import (
	"Bank-INA/config"
	"Bank-INA/domain"

	"github.com/gin-gonic/gin"
)

var (
	cfg config.Config
)

func response(ctx *gin.Context, status int, data interface{}) {
	ctx.JSON(status, data)
}

func Failed(ctx *gin.Context, data domain.ErrorData) {
	if cfg.Debug == "prod" {
		data.Data = nil
	}

	response(ctx, 400, data)
}

func Unauth(ctx *gin.Context, data domain.ErrorData) {

	response(ctx, 401, data)
}

func Success(ctx *gin.Context, data interface{}) {
	dataRes := domain.Response{
		ResponseCode: "0000",
		Message:      "success",
		Data:         data,
	}

	response(ctx, 200, dataRes)
}

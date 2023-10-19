package router

import "github.com/gin-gonic/gin"

func userRouter(api *gin.RouterGroup) {
	userControll := userRouterInit()
	api.POST("/users", userControll.Create)
	api.GET("/users", userControll.Get)
	api.GET("/users/:id", userControll.GetByID)
	api.PUT("/users/:id", userControll.Update)
	api.DELETE("/users/:id", userControll.Delete)
}

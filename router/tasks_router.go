package router

import "github.com/gin-gonic/gin"

func taskRouter(api *gin.RouterGroup) {
	taskControll := tasksRouterInit()
	api.POST("/tasks", setMiddleware().Authenticate(), taskControll.Create)
	api.GET("/tasks", setMiddleware().Authenticate(), taskControll.Get)
	api.GET("/tasks/:id", setMiddleware().Authenticate(), taskControll.GetByID)
	api.PUT("/tasks/:id", setMiddleware().Authenticate(), taskControll.Update)
	api.DELETE("/tasks/:id", setMiddleware().Authenticate(), taskControll.Delete)
}

package router

import "github.com/gin-gonic/gin"

func taskRouter(api *gin.RouterGroup) {
	taskControll := tasksRouterInit()
	api.POST("/tasks", taskControll.Create)
	api.GET("/tasks", taskControll.Get)
	api.GET("/tasks/:id", taskControll.GetByID)
	api.PUT("/tasks/:id", taskControll.Update)
	api.DELETE("/tasks/:id", taskControll.Delete)
}

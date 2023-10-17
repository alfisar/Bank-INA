package router

import "github.com/gin-gonic/gin"

func welcomeRouter(api *gin.RouterGroup) {
	welcomControll := welcomeRouterInit()
	api.GET("/", welcomControll.Hello)
}

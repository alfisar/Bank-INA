package router

import "github.com/gin-gonic/gin"

func NewRoute() *gin.Engine {
	api := gin.Default()

	v1 := api.Group("/api/v1")
	welcomeRouter(v1)

	return api
}

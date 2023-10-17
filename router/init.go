package router

import (
	_welcomeController "Bank-INA/application/welcome/controller"
	"Bank-INA/database"

	"github.com/joho/godotenv"
)

var (
	_  = godotenv.Load(".env")
	db = database.NewConnMysql()
)

func welcomeRouterInit() *_welcomeController.ControllerWelcome {
	return _welcomeController.NewWelcomeController()
}

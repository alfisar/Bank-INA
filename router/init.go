package router

import (
	_welcomeController "Bank-INA/application/welcome/controller"
	"os"

	_userController "Bank-INA/application/user/controller"
	_userMysql "Bank-INA/application/user/repository/mysql"
	_userService "Bank-INA/application/user/service"

	_tasksController "Bank-INA/application/tasks/controller"
	_tasksMysql "Bank-INA/application/tasks/repository/mysql"
	_tasksService "Bank-INA/application/tasks/service"

	"Bank-INA/internal/middleware"
	oauthhandler "Bank-INA/internal/oauthHandler"

	"Bank-INA/database"

	"github.com/joho/godotenv"
)

var (
	_  = godotenv.Load(".env")
	db = database.NewConnMysql(0)
)

func welcomeRouterInit() *_welcomeController.ControllerWelcome {
	return _welcomeController.NewWelcomeController()
}

func userRouterInit() *_userController.UserController {
	repo := _userMysql.NewUserRepository(db)
	serv := _userService.NewUserService(repo)
	controll := _userController.NewUserController(serv)
	return controll
}

func tasksRouterInit() *_tasksController.TasksController {
	repo := _tasksMysql.NewTaskRepository(db)
	serv := _tasksService.NewTasksService(repo)
	controll := _tasksController.NewTasksController(serv)
	return controll
}

func setMiddleware() *middleware.AuthenticateMiddleware {
	oauth := oauthhandler.GetOAuthHandler()
	oauth.ClientID = os.Getenv("OAUTH_CLIENT_ID")
	oauth.ClientSecret = os.Getenv("OAUTH_CLIENT_SECRET")
	oauth.TokenUrl = os.Getenv("OAUTH_TOKEN_URL")
	middleWR := middleware.NewAuthenticateMiddleware(oauth)
	return middleWR
}

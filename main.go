package main

import "Bank-INA/router"

func main() {
	route := router.NewRoute()
	route.Run("localhost:8080")
}

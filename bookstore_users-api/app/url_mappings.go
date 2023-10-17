package app

import (
	ping "bookstore_users-api/controllers/ping"
	users "bookstore_users-api/controllers/users"
)

func mapurls() {

	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
	router.PUT("/users/:user_id", users.UpdateUser)
	router.PATCH("/users/:user_id", users.UpdateUser)
	router.DELETE("/users/:user_id", users.DeleteUser)
	router.GET("/internal/users/search", users.Search)
	router.POST("/users/login", users.Login)
}

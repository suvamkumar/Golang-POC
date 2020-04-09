package app

import "crud_with_gin_gonic/controllers/user"

func mapUrls() {
	router.POST("/users", user.Create)
	router.PATCH("/users/:id", user.UpdateUser)
	router.DELETE("/users/:id", user.DeleteUser)
	router.GET("/users/:id", user.GetUser)
	router.GET("/users", user.GetAllUser)
	router.GET("/ping", user.Hello)
}

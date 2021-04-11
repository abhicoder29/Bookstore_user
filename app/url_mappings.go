package app

import (
	"github.com/Bookstore_user/controllers/ping"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", controllers.GetUsers);
	router.GET("/users/:search", controllers.SearchUsers);
	router.POST("/users", controllers.CreateUsers);
	
	// router.GET("/users/:user_id", users.Get)
	// router.PUT("/users/:user_id", users.Update)
	// router.PATCH("/users/:user_id", users.Update)
	// router.DELETE("/users/:user_id", users.Delete)
	// router.GET("/internal/users/search", users.Search)
	// router.POST("/users/login", users.Login)
}

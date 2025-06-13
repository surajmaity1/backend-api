package routes

import (
	"github.com/julienschmidt/httprouter"
	"github.com/surajmaity1/backend-api/controllers"
)

func PostRoutes(router *httprouter.Router) {
	router.POST("/post", controllers.CreatePost)
	//router.GET("/post", controllers.GetPost)
}

package routes

import (
	"github.com/julienschmidt/httprouter"
	"github.com/surajmaity1/backend-api/controllers"
)

func PostRoutes(router *httprouter.Router) {
	router.POST("/posts", controllers.CreatePost)
	router.GET("/posts/:id", controllers.GetPost)
}

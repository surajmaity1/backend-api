package routes

import (
	"backend-api/controllers"
	"github.com/julienschmidt/httprouter"
)

func PostRoutes(router *httprouter.Router) {
	router.POST("/post", controllers.HealthCheckHandler)
}

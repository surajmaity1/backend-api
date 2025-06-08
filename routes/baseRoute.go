package routes

import (
	"backend-api/controllers"
	"github.com/julienschmidt/httprouter"
)

func SetupBaseRoutes(router *httprouter.Router) {
	router.GET("/health", controllers.HealthCheckHandler)
}

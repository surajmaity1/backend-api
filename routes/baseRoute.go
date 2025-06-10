package routes

import (
	"github.com/julienschmidt/httprouter"
	"github.com/surajmaity1/backend-api/controllers"
)

func SetupBaseRoutes(router *httprouter.Router) {
	router.GET("/health", controllers.HealthCheckHandler)
}

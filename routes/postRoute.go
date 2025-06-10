package routes

import (
	"github.com/julienschmidt/httprouter"
	_ "github.com/surajmaity1/backend-api/services"
)

func PostRoutes(router *httprouter.Router) {
	//router.POST("/post", services.AddPost)
}

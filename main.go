package main

import (
	_ "github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/surajmaity1/backend-api/config"
	"github.com/surajmaity1/backend-api/routes"
)

func main() {
	logrus.Info("Server running on port " + config.AppConfig.SERVER_PORT)
	routes.Listen(":" + config.AppConfig.SERVER_PORT)

}

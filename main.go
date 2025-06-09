package main

import (
	"backend-api/config"
	"backend-api/routes"
	_ "github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Server running on port " + config.AppConfig.SERVER_PORT)
	routes.Listen(":" + config.AppConfig.SERVER_PORT)

}

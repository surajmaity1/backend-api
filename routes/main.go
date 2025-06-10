package routes

import (
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Listen(listenAddress string) {
	router := httprouter.New()
	SetupBaseRoutes(router)
	PostRoutes(router)
	err := http.ListenAndServe(listenAddress, router)
	if err != nil {
		logrus.Fatal(err)
	}
}

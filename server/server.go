package server

import (
	"cimble/server/router"
	"cimble/utilities"
	"fmt"
)

func SetupServer() {

	router := router.SetupRoutes()

	port := utilities.GetEnvironmentVariableString("PORT")
	router.Router.Run(fmt.Sprintf(":%s", port))
}

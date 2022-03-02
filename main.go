package main

import (
	"cimble/server"
	"cimble/utilities"
	"fmt"
)

func main() {

	err := utilities.LoadEnvironmentVariables()
	if err != nil {
		fmt.Printf("Unable to load env variables: %v", err.Error())
		return
	}

	_, err = utilities.ConnectToDatabase()
	if err != nil {
		fmt.Printf("Unable to connect to database %v", err.Error())
		return
	}

	server.SetupServer()
}

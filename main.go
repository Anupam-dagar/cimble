package main

import (
	"cimble/utilities"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	config, err := utilities.LoadEnvironmentVariables()
	if err != nil {
		fmt.Printf("Unable to load env variables: %v", err.Error())
		return
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})

	router.Run(fmt.Sprintf(":%s", config.PORT))
}

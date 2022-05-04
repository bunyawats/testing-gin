package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	SetupServer().Run()

}

func SetupServer() *gin.Engine {

	router := gin.Default()
	router.GET("/", IndexHandler)
	return router
}

func IndexHandler(context *gin.Context) {

	context.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}

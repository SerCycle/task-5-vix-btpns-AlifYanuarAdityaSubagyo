package main

import (
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/go-playground/validator/v10"

	"github.com/SerCycle/BTPNFinalProject/handler"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/hello", helloHandler)
	router.POST("/users/register", handler.RegisterHandler)
	router.GET("/users/login", handler.LoginHandler)

	router.Run(":8765")
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"test": "aman",
	})
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hasil": "siap",
	})
}

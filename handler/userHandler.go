package handler

import (
	"fmt"
	"net/http"

	"github.com/SerCycle/BTPNFinalProject/inputItem"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func LoginHandler(c *gin.Context) {
	email := c.Query("email")
	password := c.Query("password")

	c.JSON(http.StatusOK, gin.H{"email": email, "password": password})
}

func RegisterHandler(c *gin.Context) {
	var register inputItem.Register

	err := c.ShouldBindJSON(&register)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"username": register.Username,
		"email":    register.Email,
		"password": register.Password,
	})
}

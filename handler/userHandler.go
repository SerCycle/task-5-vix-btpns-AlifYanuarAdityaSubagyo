package handler

import (
	"fmt"
	"net/http"
	"strconv"

	// "github.com/SerCycle/BTPNFinalProject/inputItem"
	"github.com/SerCycle/BTPNFinalProject/user"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) GetUserList(c *gin.Context) {
	users, err := h.userService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var usersResponse []user.UserResponse

	for _, u := range users {
		userResponse := convertToUserResponse(u)

		usersResponse = append(usersResponse, userResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": usersResponse,
	})
}

func (h *userHandler) GetUser(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	u, err := h.userService.FindByID(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	userResponse := convertToUserResponse(u)

	c.JSON(http.StatusOK, gin.H{
		"data": userResponse,
	})
}

func (h *userHandler) LoginHandler(c *gin.Context) {
	email := c.Query("email")
	password := c.Query("password")

	c.JSON(http.StatusOK, gin.H{"email": email, "password": password})
}

func (h *userHandler) RegisterHandler(c *gin.Context) {
	var register user.Register

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

	user, err := h.userService.Create(register)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (h *userHandler) UpdateHandler(c *gin.Context) {
	var register user.Register

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

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	user, err := h.userService.Update(id, register)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	userResponse := convertToUserResponse(user)

	c.JSON(http.StatusOK, gin.H{
		"data": userResponse,
	})
}

func (h *userHandler) DelUser(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	u, err := h.userService.Delete(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	userResponse := convertToUserResponse(u)

	c.JSON(http.StatusOK, gin.H{
		"data": userResponse,
	})
}

func convertToUserResponse(u user.User) user.UserResponse {
	return user.UserResponse{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
	}
}

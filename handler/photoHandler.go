package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/SerCycle/BTPNFinalProject/photo"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type photoHandler struct {
	photoService photo.Service
}

func NewPhotoHandler(photoService photo.Service) *photoHandler {
	return &photoHandler{photoService}
}

func (h *photoHandler) GetPhotoList(c *gin.Context) {
	photos, err := h.photoService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var photosResponse []photo.PhotoResponse

	for _, p := range photos {
		photoResponse := convertToPhotoResponse(p)

		photosResponse = append(photosResponse, photoResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": photosResponse,
	})
}

func (h *photoHandler) GetPhoto(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	p, err := h.photoService.FindByID(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	photoResponse := convertToPhotoResponse(p)

	c.JSON(http.StatusOK, gin.H{
		"data": photoResponse,
	})
}

func (h *photoHandler) AddPhotoHandler(c *gin.Context) {
	var addPhoto photo.AddPhoto

	err := c.ShouldBindJSON(&addPhoto)
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

	photo, err := h.photoService.Create(addPhoto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": photo,
	})
}

func (h *photoHandler) UpdateHandler(c *gin.Context) {
	var addPhoto photo.AddPhoto

	err := c.ShouldBindJSON(&addPhoto)
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

	photo, err := h.photoService.Update(id, addPhoto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	photoResponse := convertToPhotoResponse(photo)

	c.JSON(http.StatusOK, gin.H{
		"data": photoResponse,
	})
}

func (h *photoHandler) DelPhoto(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	p, err := h.photoService.Delete(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	userResponse := convertToPhotoResponse(p)

	c.JSON(http.StatusOK, gin.H{
		"data": userResponse,
	})
}

func convertToPhotoResponse(p photo.Photo) photo.PhotoResponse {
	return photo.PhotoResponse{
		ID:       p.ID,
		Title:    p.Title,
		Caption:  p.Caption,
		PhotoUrl: p.PhotoUrl,
		UserID:   p.UserID,
	}
}

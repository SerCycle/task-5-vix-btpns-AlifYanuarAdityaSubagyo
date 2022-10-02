package main

import (
	"log"
	"net/http"

	"github.com/SerCycle/BTPNFinalProject/handler"
	"github.com/SerCycle/BTPNFinalProject/photo"
	"github.com/SerCycle/BTPNFinalProject/user"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:1202190187@tcp(127.0.0.1:3306)/btpnapigolang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connetion Error")
	}

	db.AutoMigrate(&photo.Photo{})

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	photoRepository := photo.NewRepository(db)
	photoService := photo.NewService(photoRepository)
	photoHandler := handler.NewPhotoHandler(photoService)

	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/users", userHandler.GetUserList)
	router.GET("/users/:id", userHandler.GetUser)
	router.PUT("/users/:id", userHandler.UpdateHandler)
	router.POST("/users/register", userHandler.RegisterHandler)
	router.GET("/users/login", userHandler.LoginHandler)
	router.DELETE("/users/:id", userHandler.DelUser)

	router.POST("/photos", photoHandler.AddPhotoHandler)
	router.GET("/photos", photoHandler.GetPhotoList)
	router.GET("/photos/:id", photoHandler.GetPhoto)
	router.PUT("/photos/:id", photoHandler.UpdateHandler)
	router.DELETE("/photos/:id", photoHandler.DelPhoto)

	router.Run(":8765")
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Status": "Server Aman",
	})
}

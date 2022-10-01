package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/go-playground/validator/v10"

	"github.com/SerCycle/BTPNFinalProject/handler"
	// "github.com/SerCycle/BTPNFinalProject/model"
	"github.com/SerCycle/BTPNFinalProject/photo"
	"github.com/SerCycle/BTPNFinalProject/user"

	// "github.com/SerCycle/BTPNFinalProject/repository"
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

	// var ListUser model.User

	// err = db.Debug().Where("email = ?", "Gimmick@mail.com").Find(&ListUser).Error
	// ListUser.Username = "pertamaasik"

	// if err != nil {
	// 	fmt.Println("ga ketarik")
	// }

	// err = db.Save(&ListUser).Error
	// if err != nil {
	// 	fmt.Println("ga ke update")
	// }

	// var ListUser model.User

	// err = db.Debug().Where("username = ?", "asmarajinggo").Find(&ListUser).Error
	// if err != nil {
	// 	fmt.Println("ga ketarik")
	// }
	// err = db.Delete(&ListUser).Error
	// if err != nil {
	// 	fmt.Println("ga keapus")
	// }

	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/hello", helloHandler)
	router.GET("/users", userHandler.GetUserList)
	router.GET("/users/:id", userHandler.GetUser)
	router.PUT("/users/:id", userHandler.UpdateHandler)
	router.POST("/users/register", userHandler.RegisterHandler)
	router.GET("/users/login", userHandler.LoginHandler)
	router.DELETE("/users/:id", userHandler.DelUser)

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

package main

import (
	// "fmt"

	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/go-playground/validator/v10"

	"github.com/SerCycle/BTPNFinalProject/handler"
	"github.com/SerCycle/BTPNFinalProject/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:1202190187@tcp(127.0.0.1:3306)/btpnapigolang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connetion Error")
	}

	db.AutoMigrate(&model.User{})
	// Testing CRUD

	// AddUser := model.User{}
	// AddUser.Username = "Suyanto Gimmick"
	// AddUser.Email = "Gimmick@mail.com"
	// AddUser.Password = "Testing321"

	// err = db.Create(&AddUser).Error
	// if err != nil {
	// 	fmt.Println("ada kesalahan add user")
	// }

	// var ListUser []model.User

	// err = db.Debug().Where("username = ?", "Rocky Wijaya").Find(&ListUser).Error
	// if err != nil {
	// 	fmt.Println("ada kesalahan tarik data")
	// }
	// for _, lu := range ListUser {
	// 	fmt.Println("Username: ", lu.Username)
	// 	fmt.Println("ListUser", lu)
	// }

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

	var ListUser model.User

	err = db.Debug().Where("username = ?", "asmarajinggo").Find(&ListUser).Error
	if err != nil {
		fmt.Println("ga ketarik")
	}
	err = db.Delete(&ListUser).Error
	if err != nil {
		fmt.Println("ga keapus")
	}

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

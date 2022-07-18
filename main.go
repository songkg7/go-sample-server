package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/person", func(context *gin.Context) {
		person := Person{Name: "haril", Age: 28}
		context.JSON(http.StatusOK, person)
	})

	router.POST("/person/create", func(context *gin.Context) {
		var req Person
		if err := context.BindJSON(&req); err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(req.Name, req.Age)
	})

	router.Run()
}

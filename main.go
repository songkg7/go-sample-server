package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ping() func(context *gin.Context) {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}
}

func getPerson() func(context *gin.Context) {
	return func(context *gin.Context) {
		person := Person{Name: "haril", Age: 28}
		context.JSON(http.StatusOK, person)
	}
}

func create() func(context *gin.Context) {
	return func(context *gin.Context) {
		var req Person
		if err := context.BindJSON(&req); err != nil {
			fmt.Println(err.Error())
			return
		}
		context.JSON(http.StatusCreated, req)
	}
}

func main() {
	router := gin.Default()
	router.GET("/ping", ping())
	router.GET("/person", getPerson())
	router.POST("/person/create", create())
	router.Run()
}

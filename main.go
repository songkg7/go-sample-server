package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-practice/content/dblayer"
	"go-practice/content/models"
	"go-practice/content/route"
	"net/http"
	"os"
)

func main() {
	router := gin.Default()
	pgConnString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGDATABASE"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
	)

	db := dblayer.InitDatabase(pgConnString)

	// 테이블 자동생성
	db.AutoMigrate(&models.Product{})

	var product models.Product

	router.GET("/:id", route.GetValue(db, product))
	router.POST("/create", route.AddValue(db))
	router.PATCH("/:id", route.UpdateValue(db, product))
	router.DELETE("/delete/:id", route.DeleteValue(db))
	router.GET("/ping", ping())
	router.Run()
}

func ping() func(context *gin.Context) {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}
}

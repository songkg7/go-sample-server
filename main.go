package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-practice/content"
	"go-practice/content/dblayer"
	"go-practice/content/models"
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

	router.GET("/:id", content.GetValue(db, product))
	router.POST("/create", func(context *gin.Context) {
		// 생성
		result := db.Create(models.Product{
			Code:  "D42",
			Price: 100,
		})
		context.JSON(http.StatusCreated, result.RowsAffected)
	})
	router.PATCH("/:id", func(context *gin.Context) {
		// 수정 - product 의 price 를 200 으로
		db.Model(&product).Update("Price", 200)
		// 수정 - 여러개의 필드를 수정하기
		//db.Model(&product).Updates(Product{Price: 200, Code: "F42"})
		//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	})
	router.DELETE("/delete/:id", func(context *gin.Context) {
		// 삭제 - product 삭제하기
		db.Delete("&product", 1)
	})

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

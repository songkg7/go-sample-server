package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	router := gin.Default()
	db, err := gorm.Open(postgres.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("DB 연결에 실패하였습니다.")
	}

	// 테이블 자동생성
	db.AutoMigrate(&Product{})

	// 생성
	db.Create(&Product{
		Code:  "D42",
		Price: 100,
	})

	// 읽기
	var product Product
	db.First(&product, 1)                 // primary key 기준으로 product 찾기
	db.First(&product, "code = ?", "D42") // code 가 D42 인 product 찾기

	// 수정 - product 의 price 를 200 으로
	db.Model(&product).Update("Price", 200)

	// 수정 - 여러개의 필드를 수정하기
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"})
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// 삭제 - product 삭제하기
	db.Delete("&product", 1)

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

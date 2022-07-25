package route

import (
	"github.com/gin-gonic/gin"
	"go-practice/content/dblayer"
	"go-practice/content/models"
	"net/http"
)

func GetValue(db *dblayer.DbORM, product models.Product) func(context *gin.Context) {
	return func(context *gin.Context) {
		// 읽기
		db.First(&product, 1)                 // primary key 기준으로 product 찾기
		db.First(&product, "code = ?", "D42") // code 가 D42 인 product 찾기
	}
}

func UpdateValue(db *dblayer.DbORM, product models.Product) func(context *gin.Context) {
	return func(context *gin.Context) {
		// 수정 - product 의 price 를 200 으로
		db.Model(&product).Update("Price", 200)
		// 수정 - 여러개의 필드를 수정하기
		//db.Model(&product).Updates(Product{Price: 200, Code: "F42"})
		//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	}
}

func DeleteValue(db *dblayer.DbORM) func(context *gin.Context) {
	return func(context *gin.Context) {
		// 삭제 - product 삭제하기
		db.Delete("&product", 1)
	}
}

func AddValue(db *dblayer.DbORM) func(context *gin.Context) {
	return func(context *gin.Context) {
		// 생성
		result := db.Create(models.Product{
			Code:  "D42",
			Price: 100,
		})
		context.JSON(http.StatusCreated, result.RowsAffected)
	}
}

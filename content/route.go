package content

import (
	"github.com/gin-gonic/gin"
	"go-practice/content/dblayer"
	"go-practice/content/models"
)

func GetValue(db *dblayer.DbORM, product models.Product) func(context *gin.Context) {
	return func(context *gin.Context) {
		// 읽기
		db.First(&product, 1)                 // primary key 기준으로 product 찾기
		db.First(&product, "code = ?", "D42") // code 가 D42 인 product 찾기
	}
}

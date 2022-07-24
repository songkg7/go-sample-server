package dblayer

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbORM struct {
	*gorm.DB
}

func InitDatabase(dsn string) *DbORM {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DB 연결에 실패하였습니다.")
	}
	return &DbORM{db}
}

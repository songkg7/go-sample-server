package main

import (
	"database/sql"
	"fmt"
	"github.com/cenkalti/backoff/v4"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
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
	db, err := initStore()
	if err != nil {
		log.Fatalf("failed to initialise the store: %s", err)
	}
	defer db.Close()

	router.GET("/ping", ping())
	router.GET("/person", getPerson())
	router.POST("/person/create", create())
	router.Run()
}

func initStore() (*sql.DB, error) {
	pgConnString := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGDATABASE"),
		os.Getenv("PGPASSWORD"),
	)

	var (
		db  *sql.DB
		err error
	)
	openDB := func() error {
		db, err = sql.Open("postgres", pgConnString)
		return err
	}

	err = backoff.Retry(openDB, backoff.NewExponentialBackOff())
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS message (value STRING PRIMARY KEY)"); err != nil {
		return nil, err
	}
	return db, nil
}

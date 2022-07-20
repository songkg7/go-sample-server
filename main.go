package main

import (
	"database/sql"
	"fmt"
	"github.com/cenkalti/backoff/v4"
	"github.com/cockroachdb/cockroach-go/v2/crdb"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func main() {
	router := gin.Default()
	db, err := initStore()
	if err != nil {
		log.Fatalf("failed to initialise the store: %s", err)
	}
	defer db.Close()

	router.GET("/ping", ping())
	router.GET("/message", read(db))
	router.POST("/message", create(db))

	router.Run()
}

func create(db *sql.DB) func(context *gin.Context) {
	return func(context *gin.Context) {
		m := &Message{}
		if err := context.Bind(&m); err != nil {
			context.JSON(http.StatusInternalServerError, err)
		}
		crdb.ExecuteTx(context, db, nil, func(tx *sql.Tx) error {
			_, err := tx.Exec(
				"INSERT INTO message (value) VALUES ($1) ON CONFLICT (value) DO UPDATE SET value = excluded.value",
				m.Value,
			)
			if err != nil {
				//return context.JSON(http.StatusInternalServerError, err)
				return context.Error(err)
			}
			return nil
		})
		context.JSON(http.StatusCreated, m)
	}
}

func ping() func(context *gin.Context) {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}
}

func read(db *sql.DB) func(context *gin.Context) {
	return func(context *gin.Context) {
		r, err := countRecords(db)
		if err != nil {
			context.HTML(http.StatusInternalServerError, "errorTest", err.Error())
		}
		context.JSON(http.StatusOK, r)
	}
}

func countRecords(db *sql.DB) (int, error) {
	rows, err := db.Query("SELECT count(*) FROM message")
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return 0, err
		}
		rows.Close()
	}
	return count, nil
}

type Message struct {
	Value string `json:"value"`
}

func initStore() (*sql.DB, error) {
	pgConnString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGDATABASE"),
		os.Getenv("PGUSER"),
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

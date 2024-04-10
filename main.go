package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/ping", ping())
    router.GET("/long-ping", longPing())
	router.Run()
}

func ping() func(context *gin.Context) {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}
}

func longPing() func(context *gin.Context) {
    return func(context *gin.Context) {
        context.JSON(http.StatusOK, gin.H{
            "abbreviation": "CDT",
            "client_ip": "58.234.77.75",
            "datetime": "2024-04-09T21:33:36.049571-05:00",
            "day_of_week": 2,
            "day_of_year": 100,
            "dst": true,
            "dst_from": "2024-03-10T08:00:00+00:00",
            "dst_offset": 3600,
            "dst_until": "2024-11-03T07:00:00+00:00",
            "raw_offset": -21600,
            "timezone": "America/Chicago",
            "unixtime": 1712716416,
            "utc_datetime": "2024-04-10T02:33:36.049571+00:00",
            "utc_offset": "-05:00",
            "week_number": 15,
        })
    }
}


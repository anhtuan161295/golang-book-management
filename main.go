package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	controller "example.com/openweather/controller"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{"message": "Not Found"})
	})

	controller.SearchTodayWeather(r)
	controller.CreateWeatherRecord(r)
	controller.UpdateWeatherRecord(r)
	controller.DeleteWeatherRecord(r)

	return r
}

func main() {
	r := setupRouter()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/objx"

	db "example.com/openweather/database"
	service "example.com/openweather/service"
)

type CreateWeatherRequest struct {
	CityId int
	Time   int
	Data   string
}

type UpdateWeatherRequest struct {
	Id     int
	CityId int
	Time   int
	Data   string
}

type DeleteWeatherRequest struct {
	Id int
}

type OpenWeatherData struct {
	Cod     int
	Message string
}

func SearchTodayWeather(route *gin.Engine) {
	auth := route.Group("/api/weather")
	auth.GET("/search-today-weather", func(ctx *gin.Context) {

		cityName, hasParam := ctx.GetQuery("cityName")
		if !hasParam {
			ctx.JSON(400, gin.H{
				"code":    400,
				"message": "City is invalid",
			})
			return
		}

		data := service.GetWeatherByCity(cityName)

		dataObj := objx.MustFromJSON(data)
		code := dataObj.Get("cod").Str()

		if code != "200" {
			ctx.JSON(400, gin.H{
				"code":    400,
				"message": "City not found",
			})
			return
		}

		ctx.JSON(200, gin.H{
			"code":   200,
			"status": "OK",
			"data":   dataObj,
		})
	})

}

func GetWeatherRecords(route *gin.Engine) {
	auth := route.Group("/api/weather")
	auth.GET("/records", func(ctx *gin.Context) {

		cityId, hasParam := ctx.GetQuery("cityId")
		if !hasParam {
			cityId = "-1"
		}

		from, hasParam := ctx.GetQuery("from")
		if !hasParam {
			ctx.JSON(400, gin.H{
				"code":    400,
				"message": "From parameter is missing",
			})
			return
		}

		to, hasParam := ctx.GetQuery("to")
		if !hasParam {
			ctx.JSON(400, gin.H{
				"code":    400,
				"message": "To parameter is missing",
			})
			return
		}

		page, hasParam := ctx.GetQuery("page")
		if !hasParam {
			page = "1"
		}

		size, hasParam := ctx.GetQuery("size")
		if !hasParam {
			size = "10"
		}

		parsedCityId, _ := strconv.Atoi(cityId)
		parsedPage, _ := strconv.Atoi(page)
		parsedSize, _ := strconv.Atoi(size)

		data := db.GetWeathers(parsedCityId, from, to, parsedPage, parsedSize)

		ctx.JSON(200, gin.H{
			"code":   200,
			"status": "OK",
			"data":   data,
		})
	})

}

func CreateWeatherRecord(route *gin.Engine) {
	auth := route.Group("/api/weather")
	auth.POST("/record", func(ctx *gin.Context) {

		body := CreateWeatherRequest{}
		ctx.BindJSON(&body)

		weatherId := db.CreateWeather(body.CityId, body.Time, body.Data)
		if weatherId == -1 {
			ctx.JSON(400, gin.H{
				"code":    400,
				"message": "Create weather record failed",
			})
			return
		}

		ctx.JSON(200, gin.H{
			"code":   200,
			"status": "OK",
		})
	})

}

func UpdateWeatherRecord(route *gin.Engine) {
	auth := route.Group("/api/weather")
	auth.PUT("/record/:id", func(ctx *gin.Context) {

		body := UpdateWeatherRequest{}
		ctx.BindJSON(&body)

		id := ctx.Param("id")
		parsedId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(400, gin.H{
				"code":   400,
				"status": "Id is invalid",
			})
			return
		}

		body.Id = parsedId

		queryResult := db.UpdateWeather(body.Id, body.CityId, body.Time, body.Data)
		if !queryResult {
			ctx.JSON(400, gin.H{
				"code":    400,
				"message": "Update weather record failed",
			})
			return
		}

		ctx.JSON(200, gin.H{
			"code":   200,
			"status": "OK",
		})
	})

}

func DeleteWeatherRecord(route *gin.Engine) {
	auth := route.Group("/api/weather")
	auth.DELETE("/record/:id", func(ctx *gin.Context) {

		id := ctx.Param("id")
		parsedId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(400, gin.H{
				"code":   400,
				"status": "Id is invalid",
			})
			return
		}

		queryResult := db.DeleteWeather(parsedId)
		if !queryResult {
			ctx.JSON(400, gin.H{
				"code":    400,
				"message": "Delete weather record failed",
			})
			return
		}

		ctx.JSON(200, gin.H{
			"code":   200,
			"status": "OK",
		})
	})

}

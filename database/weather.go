package database

import (
	"time"

	"gorm.io/gorm"
)

type Weather struct {
	Id      int    `json:"id"`
	City_Id int    `json:"cityId"`
	Time    int    `json:"time"`
	Data    string `json:"data"`
}

// override table name
type Tabler interface {
	TableName() string
}

func (Weather) TableName() string {
	return "weather"
}

func GetWeathers(cityId int, from string, to string, page int, size int) (result []Weather) {
	conn := GetDbConnection()

	pageIndex := page - 1
	offset := pageIndex * size
	limit := size

	parsedFrom, _ := time.Parse("2006-01-02", from)
	parsedTo, _ := time.Parse("2006-01-02", to)
	fromUnix := parsedFrom.Unix()
	toUnix := parsedTo.Unix()

	var weathers []Weather
	var query *gorm.DB

	if cityId == -1 {
		query = conn.Where("time >= ?", fromUnix).Where("time <= ?", toUnix).Offset(offset).Limit(limit)
	} else {
		query = conn.Where("time >= ?", fromUnix).Where("time <= ?", toUnix).Where("city_id = ?", cityId).Offset(offset).Limit(limit)
	}

	// execute query
	query.Find(&weathers)

	return weathers
}

func CreateWeather(cityId int, time int, data string) (id int) {
	conn := GetDbConnection()

	weather := Weather{City_Id: cityId, Time: time, Data: data}
	result := conn.Create(&weather)
	if result.Error != nil {
		return -1
	}

	return weather.Id
}

func UpdateWeather(id int, cityId int, time int, data string) (res bool) {
	conn := GetDbConnection()

	weather := Weather{Id: id, City_Id: cityId, Time: time, Data: data}
	result := conn.Save(&weather)
	return result.Error == nil
}

func DeleteWeather(id int) (res bool) {
	conn := GetDbConnection()

	weather := Weather{Id: id}
	result := conn.Delete(&weather)
	return result.Error == nil
}

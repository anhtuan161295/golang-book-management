package database

type Weather struct {
	Id      int
	City_Id int
	Time    int
	Data    string
}

const TABLE_NAME = "weather"

func CreateWeather(cityId int, time int, data string) (id int) {
	conn := GetDbConnection()

	weather := Weather{City_Id: cityId, Time: time, Data: data}
	result := conn.Table(TABLE_NAME).Create(&weather)
	if result.Error != nil {
		return -1
	}

	return weather.Id
}

func UpdateWeather(id int, cityId int, time int, data string) (res bool) {
	conn := GetDbConnection()

	weather := Weather{Id: id, City_Id: cityId, Time: time, Data: data}
	result := conn.Table(TABLE_NAME).Save(&weather)
	return result.Error == nil
}

func DeleteWeather(id int) (res bool) {
	conn := GetDbConnection()

	weather := Weather{Id: id}
	result := conn.Table(TABLE_NAME).Delete(&weather)
	return result.Error == nil
}

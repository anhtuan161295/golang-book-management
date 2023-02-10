package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Configurations exported
type Configuration struct {
	Server      ServerConfig
	Database    DatabaseConfig
	OpenWeather OpenWeatherConfig
}

type ServerConfig struct {
	Port int
}

type DatabaseConfig struct {
	Url          string
	Port         string
	Username     string
	Password     string
	DatabaseName string
}

type OpenWeatherConfig struct {
	Key string
}

var MyConfiguration Configuration
var MyDatabaseConfig DatabaseConfig
var MyServerConfig ServerConfig
var MyOpenWeatherConfig OpenWeatherConfig

func getConfig() Configuration {
	// set viper configuration to read config from file
	viper.SetConfigName("dev") // name of config file (without extension)
	viper.SetConfigType("yml") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")   // optionally look for config in the working directory
	viper.AutomaticEnv()

	// check config
	var configuration Configuration

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading database config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	return configuration
}

func init() {

	MyConfiguration = getConfig()
	MyServerConfig = MyConfiguration.Server
	MyDatabaseConfig = MyConfiguration.Database
	MyOpenWeatherConfig = MyConfiguration.OpenWeather

	fmt.Println("Reading variables using the model...")
	fmt.Println("Server Port is\t", MyServerConfig.Port)

	// Reading variables using the model
	fmt.Println("Database Url is\t", MyDatabaseConfig.Url)
	fmt.Println("Database Username is\t", MyDatabaseConfig.Username)
	fmt.Println("Database Password is\t", MyDatabaseConfig.Password)

}

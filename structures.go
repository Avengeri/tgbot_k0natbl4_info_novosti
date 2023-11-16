package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

// Переменная через которую идет управление и настройка бота
var gBot *tgbotapi.BotAPI
var config Config

// Структура конфига для подключения к БД
type DbConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

// Структура, чтобы создать вложенную структуру (чтобы был один конфиг для всего)
type Config struct {
	Storage DbConfig `yaml:"storage"`
}

// Основная структура юзера, куда будут лететь обновления из ТГ
type User struct {
	id          int64
	chatId      int64
	userName    string
	money       int64
	hasLocation bool
	latitude    float64
	longitude   float64
}

//Структура для информации о погоде

type WeatherResponse struct {
	Fact struct {
		Temperature int64  `json:"temp"`
		FeelsLike   int64  `json:"feels_like"`
		Icon        string `json:"icon"`
		Condition   string `json:"condition"`
		WindSpeed   int64  `json:"wind_speed"`
		Pressure    int64  `json:"pressure_mm"`
		Humidity    int64  `json:"humidity"` //Влажность воздуха в %
	} `json:"fact"`
	Info struct {
		Url string `json:"url"`
	} `json:"info"`
	GeoObject struct {
		District struct {
			NameDistrict string `json:"name"`
		} `json:"district"`
		Locality struct {
			NameLocality string `json:"name"`
		} `json:"locality"`
		Province struct {
			NameProvince string `json:"name"`
		} `json:"province"`
		Country struct {
			NameCountry string `json:"name"`
		} `json:"country"`
	} `json:"geo_object"`
}

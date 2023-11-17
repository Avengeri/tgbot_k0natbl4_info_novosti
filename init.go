package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

// Инициализация бота через токен
func initTgBot(envFilePath string) {
	err := godotenv.Load(envFilePath)
	if err != nil {
		log.Fatal("Не удалось загрузить переменную окружения")
	}

	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		log.Fatal("Не удалось прочитать токен")
	} else {
		fmt.Printf("Значение токена: %s\n", botToken)
	}
	gBot, err = tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}
	gBot.Debug = true

	log.Printf("Имя бота: %s\n", gBot.Self.UserName)
}

// Инициализация PostgresQL
func initStorage(configFilePath string) (*Config, error) {

	configFile, err := os.Open(configFilePath)
	if err != nil {
		log.Fatal("Не удалось открыть конфигурационный файл базы данных")
		return nil, err
	}
	defer configFile.Close()

	decoder := yaml.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}
	log.Println("Декодирование прошло успешно, конфиг загружен")
	return &config, nil
}

// Инициализация Yandex погоды
func initYandexWeather(envFilePath string) (yandexAPIWeather string, err error) {
	err = godotenv.Load(envFilePath)
	if err != nil {
		log.Printf("Не удалось прочитать файл переменной окружения yandex_weather_api")
	}

	yandexAPIWeather = os.Getenv("API_YANDEX_WEATHER")
	if yandexAPIWeather == "" {
		fmt.Printf("Не удалось прочитать yandex_api")
	} else {
		fmt.Printf("Значение yandex_api: %s\n", yandexAPIWeather)
	}
	return yandexAPIWeather, nil
}

// Инициализация Yandex расписания
func initYandexSchedule(envFilePath string) (yandexAPISchedule string, err error) {
	err = godotenv.Load(envFilePath)
	if err != nil {
		log.Printf("Не удалось прочитать файл переменной окружения yandex_weather_api")
	}
	yandexAPISchedule = os.Getenv("API_YANDEX_SCHEDULE")
	if yandexAPISchedule == "" {
		fmt.Printf("Не удалось прочитать yandex_api")
	} else {
		fmt.Printf("Значение yandex_api: %s\n", yandexAPISchedule)
	}
	return yandexAPISchedule, nil
}

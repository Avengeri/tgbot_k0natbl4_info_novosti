package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Обрабатывает запрос и возвращает ответ в структуру
func (w *WeatherResponse) handlerWeather(u *User) error {

	yandexApi, err := initYandexWeather("./go.env")
	if err != nil {
		log.Println("Ошибка инициализации yandex_api")
	}

	url := fmt.Sprintf("https://api.weather.yandex.ru/v2/forecast?lat=59.955485&lon=30.295722&lang=ru_RU&limit=1")

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("X-Yandex-API-Key", yandexApi)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Не удалось выполнить запрос к Yandex")
	}
	defer resp.Body.Close() //Как-то обработать ошибку

	err = json.NewDecoder(resp.Body).Decode(&w)
	if err != nil {
		log.Println("Ошибка декодирования json файла")
	}
	return nil
}

func (w *WeatherResponse) handlerWeatherGeo(u *User) error {

	yandexApi, err := initYandexWeather("./go.env")
	if err != nil {
		log.Println("Ошибка инициализации yandex_api")
	}

	url := fmt.Sprintf("https://api.weather.yandex.ru/v2/forecast?lat=%f&lon=%f&lang=ru_RU&limit=1", u.latitude, u.longitude)

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("X-Yandex-API-Key", yandexApi)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Не удалось выполнить запрос к Yandex")
	}
	defer resp.Body.Close() //Как-то обработать ошибку

	err = json.NewDecoder(resp.Body).Decode(&w)
	if err != nil {
		log.Println("Ошибка декодирования json файла")
	}
	return nil
}

func (s *Suburban) handlerSuburban(url string) error {

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Ошибка получения гет запроса")
	}
	defer resp.Body.Close() // как-то обработать ошибку

	var response SuburbanResponse
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&response); err != nil {
		log.Println("Ошибка декодирования JSON:", err)
	}

	Schedule = response.Segments
	return nil
}

//func (u *User) handleMessage(update tgbotapi.Update) *User {
//	if update.Message != nil {
//		userID := update.Message.From.ID
//		if update.Message.Location != nil {
//			updateLatitude := update.Message.Location.Latitude
//			updateLongitude := update.Message.Location.Longitude
//
//			u.chatId = int64(userID)
//			u.latitude = updateLatitude
//			u.longitude = updateLongitude
//			u.hasLocation = true
//
//		} else {
//			if !u.hasLocation {
//				msg := tgbotapi.NewMessage(int64(userID), "Пожалуйста, разрешите геопозицию в настройках Вашего устройства.")
//				_, err := gBot.Send(msg)
//				if err != nil {
//					log.Printf("Не удалось отправить сообщение")
//				}
//			}
//		}
//	}
//	return u
//}

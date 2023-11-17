package main

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"

	"tgbot_k0natbl4_info_novosti/constans"
)

//func (u *User) updateLocation(location *tgbotapi.Location) {
//	u.location = location
//}

// Задержка
func delay(seconds uint16) {
	time.Sleep(time.Second * time.Duration(seconds))
}

// Проверка Callback запроса
func isCallBackQuery(update *tgbotapi.Update) bool {
	return update.CallbackQuery != nil && update.CallbackQuery.Data != ""
}

// Проверка геопозиционного запроса
func isUseGeo(update *tgbotapi.Update) bool {
	return update.Message != nil && update.Message.Location != nil
}

// Обновляет информацию о пользователе из обновления
func userUpdate(update tgbotapi.Update) *User {
	userNameUpdate := update.Message.From.UserName
	chatIdUpdate := update.Message.Chat.ID

	log.Printf("Имя пользователя: %s\n Chat ID: %d\n", userNameUpdate, chatIdUpdate)

	user := &User{
		chatId:   chatIdUpdate,
		userName: userNameUpdate,
		money:    0,
	}
	return user
}

func main() {
	envFilePath := "./go.env"
	initTgBot(envFilePath)

	configFilePath := "config/config.yaml"
	pgxConf, err := config.connectStorage(configFilePath)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}

	ctx := context.Background()
	pool, err := pgxpool.NewWithConfig(ctx, pgxConf)

	if err != nil {
		log.Fatalf("Не удалось создать пул соединений: %v", err)
	}
	defer pool.Close()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = constans.UPDATE_CONFIG_TIMEOUT

	updates, err := gBot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal("Не удалось получить обновление из канала")
	}

	for update := range updates {
		switch {

		case update.Message != nil && update.Message.IsCommand():

			switch update.Message.Command() {
			case "start":
				//Вытаскивает данные из обновления и создает структуру юзера
				user := userUpdate(update)
				//Проверяет есть ли юзер в БД, если нет -создает
				user.checkUser(ctx, pool)
				//Вытаскивает юзера из БД
				_, err = fetchUser(ctx, pool, update)
				if err != nil {
					log.Printf("Ошибка при получении информации о пользователе: %v", err)
					return
				}
				//Блок обработки
				user.showKeyboardMenuWithButton("А вот и кнопочка")

			}

		case update.Message != nil:
			user, err := fetchUser(ctx, pool, update)
			if err != nil {
				log.Printf("Ошибка при получении информации о пользователе: %v", err)
				return
			}
			log.Println(user)

			if update.Message.Text != "" {
				user.processKeyboardForText(&update, WeatherResponse{}, Suburban{})

			} else {
				user.latitude = update.Message.Location.Latitude
				user.longitude = update.Message.Location.Longitude

				_, err := pool.Exec(
					ctx,
					"UPDATE tg_users SET latitude = $1, longitude = $2 WHERE chat_id = $3",
					user.latitude,
					user.longitude,
					user.chatId,
				)

				if err != nil {
					log.Printf("Не удалось выполнить обновление в базе данных: %v", err)

				}

				if user.latitude != 0 && user.longitude != 0 {
					w := WeatherResponse{}
					w.sendTemperatureToUserGeo(user.chatId, user)
				} else {
					msg := tgbotapi.NewMessage(user.chatId, "Не удалось получить данные о Вашей геопозиции. Пожалуйста проверьте настройки приложения")
					_, err := gBot.Send(msg)
					if err != nil {
						log.Printf("Не удалось отправить сообщение")
					}
				}
			}

		case isCallBackQuery(&update):
			user, err := fetchСallBackUser(ctx, pool, update)
			if err != nil {
				log.Printf("Ошибка при получении информации о пользователе: %v", err)
				return
			}
			user.processUpdatingCallBack(&update, ctx, pool)

			//case isUseGeo(&update):
			//	user, err := fetchСallBackUser(ctx, pool, update)
			//	if err != nil {
			//		log.Printf("Ошибка при получении информации о пользователе: %v", err)
			//		return
			//	}
			//	msg := tgbotapi.NewMessage(user.chatId, "Геопозиция получена, можете попробовать посмотреть температуру рядом с собой")
			//	_, err = gBot.Send(msg)
			//	if err != nil {
			//		log.Printf("Не удалось отправить сообщение")
			//	}
			//	user.processKeyboardForText(&update, WeatherResponse{})
		}

	}
}

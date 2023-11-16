package main //Обработка команд
//import (
//	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
//	"log"
//	"tgbot_k0natbl4_info_novosti/constans"
//)
//
//func lalala() {
//	if update.Message != nil && update.Message.IsCommand() {
//
//		switch update.Message.Command() {
//		case "start":
//			user := userUpdate(update)
//			user.checkUser(ctx, pool)
//			err = user.fetchUser(ctx, pool)
//			if err != nil {
//				log.Printf("Ошибка при получении информации о пользователе: %v", err)
//				return
//			}
//			user.showKeyboardMenuWithButton("А вот и кнопочка")
//		}
//		//Обработка входящего текста
//	} else if update.Message != nil {
//		userID := update.Message.From.ID
//		var user *User
//
//		switch update.Message.Text {
//		case constans.BUTTON_REPLY_TEXT_WEATHER:
//			if update.Message.Location != nil {
//				updateLatitude := update.Message.Location.Latitude
//				updateLongitude := update.Message.Location.Longitude
//
//				user = &User{
//					chatId:    int64(userID),
//					latitude:  updateLatitude,
//					longitude: updateLongitude,
//				}
//
//			} else {
//				{
//					msg := tgbotapi.NewMessage(int64(userID), "Пожалуйста, разрешите геопозицию в настройках Вашего устройства.")
//					_, err := gBot.Send(msg)
//					if err != nil {
//						log.Printf("Не удалось отправить сообщение")
//					}
//				}
//			}
//		}
//		user.checkUser(ctx, pool)
//		err = user.fetchUser(ctx, pool)
//		if err != nil {
//			log.Printf("Ошибка при получении информации о пользователе: %v", err)
//			return
//		}
//		var w WeatherResponse
//		user.processKeyboardForText(&update, w)
//
//		//Обработка callback запросов
//	} else if isCallBackQuery(&update) {
//		userID := update.CallbackQuery.From.ID
//
//		user := &User{
//			chatId: int64(userID),
//		}
//		user.checkUser(ctx, pool)
//		err = user.fetchUser(ctx, pool)
//		if err != nil {
//			log.Printf("Ошибка при получении информации о пользователе: %v", err)
//			return
//		}
//		var w WeatherResponse
//		user.processUpdatingCallBack(&update, ctx, pool, w)
//	}
//}

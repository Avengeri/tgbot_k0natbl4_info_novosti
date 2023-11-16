package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"tgbot_k0natbl4_info_novosti/constans"
)

// Отправляет картинку
func (u *User) sendImage() tgbotapi.ReplyKeyboardMarkup {
	msg := tgbotapi.NewMessage(u.chatId, "Как тебе такое?")
	_, err := gBot.Send(msg)
	if err != nil {
		log.Printf("Не удалось отправить сообщение")
	}
	photo := tgbotapi.NewPhotoUpload(u.chatId, "/home/u-andrey/Загрузки/photo_2023-11-10_13-31-18.jpg")
	_, err = gBot.Send(photo)
	if err != nil {
		log.Printf("Не удалось отправить сообщение")
	}
	return tgbotapi.NewReplyKeyboard()
}

// Создает ряд кнопок в reply клавиатуре
func newKeyboardRow(buttonTexts ...string) []tgbotapi.KeyboardButton {
	var buttons []tgbotapi.KeyboardButton

	for _, text := range buttonTexts {
		buttons = append(buttons, tgbotapi.NewKeyboardButton(text))
	}
	return buttons
}

// Создает строку из Inline кнопок URL
func newInlineKeyboardWithUrl(buttonText, buttonUrl string) []tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonURL(buttonText, buttonUrl))
}

// Создает строку из Inline кнопок
func newInlineKeyboard(buttonText, buttonCode string) []tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(buttonText, buttonCode))
}

// Добавляет кнопку (прикреплять к сообщению)
func (u *User) buttonToMessage(msg, buttonText, buttonCode string) {
	inlineButton := newInlineKeyboard(buttonText, buttonCode)
	inlineKeeyboard := tgbotapi.NewInlineKeyboardMarkup(inlineButton)

	msg2 := tgbotapi.NewMessage(u.chatId, msg)
	msg2.ReplyMarkup = inlineKeeyboard
	_, err := gBot.Send(msg2)
	if err != nil {
		log.Printf("Не удалось отправить сообщение")
	}
}

// Кнопочка
func (u *User) buttonLalala(msgInline string) tgbotapi.ReplyKeyboardMarkup {
	inlineButton := newInlineKeyboardWithUrl(constans.BUTTON_INLINE_BUTTON_MAIN_MENU, constans.BUTTON_INLINE_BUTTON_MAIN_MENU_URL)
	secondInlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(inlineButton)

	msg2 := tgbotapi.NewMessage(u.chatId, msgInline)
	msg2.ReplyMarkup = secondInlineKeyboard
	_, err := gBot.Send(msg2)
	if err != nil {
		log.Println("Ошибка при отправке сообщения:", err)
	}
	return tgbotapi.NewReplyKeyboard()
}

// Показывает информацию об аккаунте
//func (u *User) showAccountManagement() tgbotapi.ReplyKeyboardMarkup {
//	replyRow := newKeyboardRow(constans.BUTTON_REPLY_TEXT_ACCOUNT_MANAGEMENT_SHOW_MONEY)
//	replyRow2 := newKeyboardRow(constans.BUTTON_REPLY_TEXT_ACCOUNT_MANAGEMENT_BUY_MONEY)
//	replyRow4 := newKeyboardRow(constans.BUTTON_REPLY_TEXT_ACCOUNT_MANAGEMENT_SPEND_MONEY)
//	replyRow_3 := newKeyboardRow(constans.BUTTON_REPLY_TEXT_BACK)
//	replyKeyboard := tgbotapi.NewReplyKeyboard(replyRow2, replyRow, replyRow4, replyRow_3)
//	msg := tgbotapi.NewMessage(u.chatId, "")
//	msg.ReplyMarkup = replyKeyboard
//	_, err := gBot.Send(msg)
//	if err != nil {
//		log.Printf("Не удалось отправить сообщение")
//	}
//	return replyKeyboard
//}

// Показать меню с новостями
func (u *User) showNews() {
	msg := tgbotapi.NewMessage(u.chatId, "Выбери новости какие хочешь")

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		newInlineKeyboard(constans.BUTTON_INLINE_TEXT_NEWS_GAME, constans.BUTTON_INLINE_CODE_NEWS_GAME),
		newInlineKeyboard(constans.BUTTON_INLINE_TEXT_NEWS_SERIAL, constans.BUTTON_INLINE_CODE_NEWS_SERIAL),
		newInlineKeyboard(constans.BUTTON_INLINE_TEXT_NEWS_ANIMALS, constans.BUTTON_INLINE_CODE_NEWS_ANIMALS),
	)
	_, err := gBot.Send(msg)
	if err != nil {
		log.Printf("Не удалось отправить сообщение")
	}
}

// Показывает услуги
func (u *User) showServices() {
	msg := tgbotapi.NewMessage(u.chatId, "Какие услуги будем заказывать?")

	replyRow := newKeyboardRow(constans.BUTTON_REPLY_TEXT_NEWS)
	replyRow_2 := newKeyboardRow(constans.BUTTON_REPLY_TEXT_WEATHER)
	replyRow_3 := newKeyboardRow(constans.BUTTON_REPLY_TEXT_DOLLAR)
	replyRow_4 := newKeyboardRow(constans.BUTTON_REPLY_TEXT_BACK)

	replyKeyboard := tgbotapi.NewReplyKeyboard(replyRow, replyRow_2, replyRow_3, replyRow_4)

	msg.ReplyMarkup = replyKeyboard

	_, err := gBot.Send(msg)
	if err != nil {
		log.Printf("Не удалось отправить сообщение")
	}
}

// Показывает баланс
func (u *User) showBalance() {
	msg := fmt.Sprintf("%s, сейчас у тебя %d %s", u.userName, u.money, constans.EMOJI_COIN)
	resultMsg := tgbotapi.NewMessage(u.chatId, msg)
	_, err := gBot.Send(resultMsg)
	if err != nil {
		log.Printf("Не удалось отправить сообщение")
	}
}

// Меню управления аккаунтом
func (u *User) showAccountManagement() {
	msg := tgbotapi.NewMessage(u.chatId, "Настрой свой аккаунт")

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		newInlineKeyboard(constans.BUTTON_INLINE_TEXT_ACCOUNT_MANAGEMENT_BUY_MONEY, constans.BUTTON_INLINE_CODE_ACCOUNT_MANAGEMENT_BUY_MONEY),
		newInlineKeyboard(constans.BUTTON_INLINE_TEXT_ACCOUNT_MANAGEMENT_SPEND_MONEY, constans.BUTTON_INLINE_CODE_ACCOUNT_MANAGEMENT_SPEND_MONEY),
		newInlineKeyboard(constans.BUTTON_INLINE_TEXT_ACCOUNT_MANAGEMENT_SHOW_MONEY, constans.BUTTON_INLINE_CODE_ACCOUNT_MANAGEMENT_SHOW_MONEY),
	)
	_, err := gBot.Send(msg)
	if err != nil {
		log.Printf("Не удалось отправить сообщение")
	}
}

// Показывает клавиатуру с выбором погоды
func (u *User) showWeatherChoose() {
	msg := tgbotapi.NewMessage(u.chatId, "Где будем смотреть погоду?")

	replyRow := newKeyboardRow(constans.BUTTON_REPLY_TEXT_WEATHER_SAINT_PETERSBURG)
	replyRow_2 := newKeyboardRow(constans.BUTTON_REPLY_TEXT_WEATHER_GEO)
	replyRow_3 := newKeyboardRow(constans.BUTTON_REPLY_TEXT_ALLOW_GEO)
	replyRow_4 := newKeyboardRow(constans.BUTTON_REPLY_TEXT_BACK)
	replyKeyboard := tgbotapi.NewReplyKeyboard(replyRow, replyRow_2, replyRow_3, replyRow_4)

	msg.ReplyMarkup = replyKeyboard

	_, err := gBot.Send(msg)
	if err != nil {
		log.Printf("Не удалось отправить сообщение")
	}
}

// Разрешить использовать ГЕО
func (u *User) allowGeoButton(update *tgbotapi.Update) {
	locationButton := tgbotapi.NewKeyboardButton(constans.BUTTON_REPLY_TEXT_SEND_GEO)
	locationButton.RequestLocation = true

	menuButton := tgbotapi.NewKeyboardButton(constans.BUTTON_REPLY_TEXT_BACK)

	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(locationButton),
		tgbotapi.NewKeyboardButtonRow(menuButton),
	)

	msg := tgbotapi.NewMessage(u.chatId, "Пожалуйста поделитесь геопозицией")
	msg.ReplyMarkup = keyboard

	_, err := gBot.Send(msg)
	if err != nil {
		log.Printf("Не удалось отправить сообщение")

	}

}

// Показывает клавиатуру главного меню и присылает кнопку
func (u *User) showKeyboardMenuWithButton(msgInline string) tgbotapi.ReplyKeyboardMarkup {
	replyRow := newKeyboardRow(constans.BUTTON_REPLY_TEXT_SERVICES)
	replyRow_2 := newKeyboardRow(constans.BUTTON_REPLY_TEXT_ACCOUNT_MANAGEMENT)
	replyRow_3 := newKeyboardRow(constans.BUTTON_REPLY_TEXT_ABOUT_ME)
	replyRow_4 := newKeyboardRow(constans.BUTTON_REPLY_TEXT_WANT_BUTTON, constans.BUTTON_REPLY_TEXT_WANT_PUCTURE)
	firstReplyKeyboard := tgbotapi.NewReplyKeyboard(replyRow, replyRow_4, replyRow_2, replyRow_3)
	msg1 := fmt.Sprintf("Привет %s! Функционал этого бота чисто побаловаться. Я хочу реализовать два вида клавиатур. По-этому я вынужден отправить тебе кнопку.", u.userName)
	msgResult := tgbotapi.NewMessage(u.chatId, msg1)
	msgResult.ReplyMarkup = firstReplyKeyboard
	_, err := gBot.Send(msgResult)
	if err != nil {
		log.Println("Ошибка при отправке сообщения:", err)
	}
	delay(1)

	u.buttonLalala(msgInline)

	return firstReplyKeyboard
}

// Показывает клавиатуру главного меню, но не присылает кнопку
func (u *User) showKeyboardMenuWithoutButton() tgbotapi.ReplyKeyboardMarkup {

	replyRow := newKeyboardRow(constans.BUTTON_REPLY_TEXT_SERVICES)
	replyRow_2 := newKeyboardRow(constans.BUTTON_REPLY_TEXT_ACCOUNT_MANAGEMENT)
	replyRow_3 := newKeyboardRow(constans.BUTTON_REPLY_TEXT_ABOUT_ME)
	replyRow_4 := newKeyboardRow(constans.BUTTON_REPLY_TEXT_WANT_BUTTON, constans.BUTTON_REPLY_TEXT_WANT_PUCTURE)

	firstReplyKeyboard := tgbotapi.NewReplyKeyboard(replyRow, replyRow_4, replyRow_2, replyRow_3)

	msgResult := tgbotapi.NewMessage(u.chatId, "Выберете действие")
	msgResult.ReplyMarkup = firstReplyKeyboard
	_, err := gBot.Send(msgResult)
	if err != nil {
		log.Println("Ошибка при отправке сообщения:", err)
	}
	return firstReplyKeyboard
}

// Выводит информацию о пользователе
func (u *User) showInfoAboutMe() tgbotapi.ReplyKeyboardMarkup {
	messageText := fmt.Sprintf("Мой UserName: @%s\nМой ID: %d", u.userName, u.chatId)
	msg := tgbotapi.NewMessage(u.chatId, messageText)
	_, err := gBot.Send(msg)
	if err != nil {
		log.Printf("Не удалось отправить сообщение")
	}
	return tgbotapi.NewReplyKeyboard()
}

// Отправляет погоду пользователю
func (w WeatherResponse) sendTemperatureToUser(chatID int64, u *User) {

	err := w.handlerWeather(u)
	if err != nil {
		fmt.Println("Не удалось получить значение температуры")
	} else {
		plusWeather := ""
		if w.Fact.Temperature > 0 {
			plusWeather = "+"
		}

		plusFeelsLike := ""
		if w.Fact.FeelsLike > 0 {
			plusFeelsLike = "+"
		}

		weather := w.Fact.Condition
		switch weather {
		case "wet-snow":
			weather = "дождь со снегом"
		}

		msg := fmt.Sprintf("Погода в Питере: %s%d°\nОщущается как: %s%d°\nСейчас: %s\nСкорость ветра: %d м/с\nДавление: %d мм рт. ст.\nВлажность воздуха: %d%%",
			plusWeather, w.Fact.Temperature, plusFeelsLike, w.Fact.FeelsLike, weather, w.Fact.WindSpeed, w.Fact.Pressure, w.Fact.Humidity)

		resultMessage := tgbotapi.NewMessage(chatID, msg)
		_, err = gBot.Send(resultMessage)
		if err != nil {
			log.Printf("Не удалось отправить сообщение")
		}
	}

}

// Отправляет погоду пользователю по гео
func (w WeatherResponse) sendTemperatureToUserGeo(chatID int64, u *User) {

	if u.latitude > 0 {

		err := w.handlerWeatherGeo(u)
		if err != nil {
			fmt.Println("Не удалось получить значение температуры")
		} else {
			plusWeather := ""
			if w.Fact.Temperature > 0 {
				plusWeather = "+"
			}

			plusFeelsLike := ""
			if w.Fact.FeelsLike > 0 {
				plusFeelsLike = "+"
			}

			weather := w.Fact.Condition
			switch weather {
			case "wet-snow":
				weather = "дождь со снегом"
			}

			msg := fmt.Sprintf("Ваше местоположение:%s, %s, %s, %s\nПогода: %s%d°\nОщущается как: %s%d°\nСейчас: %s\nСкорость ветра: %d м/с\nДавление: %d мм рт. ст.\nВлажность воздуха: %d%%",
				w.GeoObject.District.NameDistrict, w.GeoObject.Locality.NameLocality, w.GeoObject.Country.NameCountry, w.GeoObject.Province.NameProvince, plusWeather, w.Fact.Temperature, plusFeelsLike, w.Fact.FeelsLike, weather, w.Fact.WindSpeed, w.Fact.Pressure, w.Fact.Humidity)

			resultMessage := tgbotapi.NewMessage(chatID, msg)
			_, err = gBot.Send(resultMessage)
			if err != nil {
				log.Printf("Не удалось отправить сообщение")
			}
		}

	} else {
		u.showAllowGeo()
	}

}

func (u *User) showAllowGeo() {
	msg := tgbotapi.NewMessage(u.chatId, "Пожалуйста сначала разрешите использовать свою геопозицию (кнопка внизу). Если это не сработало, разрешено ли использовать приложению геопозицию")

	_, err := gBot.Send(msg)
	if err != nil {
		log.Printf("Не удалось отправить сообщение")
	}
}

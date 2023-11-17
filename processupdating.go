package main

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"tgbot_k0natbl4_info_novosti/constans"
	"time"
)

// Обработчик текстовых запросов + генератор новой клавиатуры + отправляет
func (u *User) processKeyboardForText(update *tgbotapi.Update, w WeatherResponse, s Suburban) {
	if update.Message == nil {
		return
	}
	text := update.Message.Text
	switch text {

	case constans.BUTTON_REPLY_TEXT_SERVICES:
		u.showServices()
	case constans.BUTTON_REPLY_TEXT_BACK_TO_SERVICES:
		u.showServices()
	case constans.BUTTON_REPLY_TEXT_ABOUT_ME:
		u.showInfoAboutMe()
	case constans.BUTTON_REPLY_TEXT_WANT_BUTTON:
		u.buttonLalala("На вот держи уже")
	case constans.BUTTON_REPLY_TEXT_WANT_PUCTURE:
		u.sendImage()
	case constans.BUTTON_REPLY_TEXT_ACCOUNT_MANAGEMENT:
		u.showAccountManagement()
	case constans.BUTTON_REPLY_TEXT_WEATHER:
		u.showWeatherChoose()
	case constans.BUTTON_REPLY_TEXT_WEATHER_SAINT_PETERSBURG:
		w.sendTemperatureToUser(u.chatId, u)
	case constans.BUTTON_REPLY_TEXT_WEATHER_GEO:
		u.allowGeoButton(update)
	case constans.BUTTON_REPLY_TEXT_NEWS:
		u.showNews()
	case constans.BUTTON_REPLY_TEXT_BACK_TO_MENU:
		u.showKeyboardMenuWithoutButton()
	case constans.BUTTON_REPLY_TEXT_BACK_TO_WEATHER:
		u.showWeatherChoose()
	case constans.BUTTON_REPLY_TEXT_SCHEDULE:
		u.showSchedule()
	case constans.BUTTON_REPLY_TEXT_SCHEDULE_TODAY_VS_SPB:
		sendScheduleToUserTodayVsSpb(u, Suburban{})
	case constans.BUTTON_REPLY_TEXT_SCHEDULE_TODAY_SPB_VS:
		sendScheduleToUserTodaySpbVs(u, Suburban{})
	default:

		tgbotapi.NewReplyKeyboard()
	}
}

// Обработчик call back запросов
func (u *User) processUpdatingCallBack(update *tgbotapi.Update, ctx context.Context, pool *pgxpool.Pool) {
	if update.CallbackQuery == nil {
		return
	}
	choiceCode := update.CallbackQuery.Data
	log.Printf("[%s],%s", time.Now(), choiceCode)
	switch choiceCode {
	case constans.BUTTON_INLINE_CODE_ACCOUNT_MANAGEMENT_BUY_MONEY:
		u.appendMoney(ctx, pool)
	case constans.BUTTON_INLINE_CODE_ACCOUNT_MANAGEMENT_SPEND_MONEY:
		u.spendMoney(ctx, pool)
	case constans.BUTTON_INLINE_CODE_ACCOUNT_MANAGEMENT_SHOW_MONEY:
		u.showBalance()

	}
}

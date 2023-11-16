package main

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"tgbot_k0natbl4_info_novosti/constans"
)

// Обновляет деньги в бд
func (u *User) storageUpdateMoney(ctx context.Context, pool *pgxpool.Pool) error {
	_, err := pool.Exec(ctx, "UPDATE tg_users SET money =$1 WHERE chat_id=$2", u.money, u.chatId)
	if err != nil {
		return err
	}
	return nil
}

// Добавляет деньги в структуру, а затем обновляет БД
func (u *User) appendMoney(ctx context.Context, pool *pgxpool.Pool) {
	u.money += 100
	if err := u.storageUpdateMoney(ctx, pool); err != nil {

	}

	msg := fmt.Sprintf("%s, спасибо, что воспользовались нашей услугой. Деньги успешно добавлены. Теперь на вашем счету: %d %s", u.userName, u.money, constans.EMOJI_COIN)
	resultMessage := tgbotapi.NewMessage(u.chatId, msg)
	_, err := gBot.Send(resultMessage)
	if err != nil {
		log.Printf("Не удалось отправить сообщение")
	}
}

// Убирает деньги из структуры, а затем обновляет БД
func (u *User) spendMoney(ctx context.Context, pool *pgxpool.Pool) {
	if u.money < 100 {
		u.buttonToMessage("У тебя не хватает денег! Пожалуйста пополни баланс.", constans.BUTTON_INLINE_TEXT_ACCOUNT_MANAGEMENT_BUY_MONEY, constans.BUTTON_INLINE_CODE_ACCOUNT_MANAGEMENT_BUY_MONEY)
	} else {
		u.money -= 100
		err := u.storageUpdateMoney(ctx, pool)
		if err != nil {
			log.Println("Не удалось списать деньги")
		}
		msg := fmt.Sprintf("%s, Вы потратили монетки. Теперь на вашем счету: %d %s", u.userName, u.money, constans.EMOJI_COIN)
		resultMessage := tgbotapi.NewMessage(u.chatId, msg)
		_, err = gBot.Send(resultMessage)
		if err != nil {
			log.Printf("Не удалось отправить сообщение")
		}
	}

}

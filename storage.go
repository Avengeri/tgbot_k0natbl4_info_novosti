package main

import (
	"context"
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

// Соединяет с бд
func (c Config) connectStorage(configFilePath string) (*pgxpool.Config, error) {
	connConf, err := initStorage(configFilePath)
	if err != nil {
		log.Fatalf("Не удалось загрузить параметры подключения из файла конфигурации: %s", err)
	}

	pgxConfig, err := pgxpool.ParseConfig("")
	if err != nil {
		log.Fatalf("Ошибка при парсинге конфигурации PGX: %v", err)
	}

	pgxConfig.ConnConfig.Host = connConf.Storage.Host
	pgxConfig.ConnConfig.Port = uint16(connConf.Storage.Port)
	pgxConfig.ConnConfig.User = connConf.Storage.User
	pgxConfig.ConnConfig.Password = connConf.Storage.Password
	pgxConfig.ConnConfig.Database = connConf.Storage.Database

	return pgxConfig, nil
}

// Проверка и создание юзера в БД
func (u *User) checkUser(ctx context.Context, pool *pgxpool.Pool) {
	err := pool.QueryRow(ctx, "SELECT id FROM tg_users WHERE chat_id=$1", u.chatId).Scan(&u.id)
	if errors.Is(err, pgx.ErrNoRows) {
		_, insertErr := pool.Exec(ctx, "INSERT INTO tg_users(chat_id, user_name, money,latitude,longitude) VALUES ($1,$2,$3,$4,$5)",
			u.chatId, u.userName, u.money, u.latitude, u.longitude)
		if insertErr != nil {
			log.Fatalf("Ошибка при создании нового пользователя: %v", insertErr)
		}
	} else if err != nil {
		log.Fatalf("Ошибка при выполнении запроса: %v", err)
	}

}

// Получение информации о пользователе из БД
func fetchUser(ctx context.Context, pool *pgxpool.Pool, update tgbotapi.Update) (*User, error) {
	user := User{}
	chatID := update.Message.From.ID
	row := pool.QueryRow(ctx, "SELECT user_name, money,latitude,longitude FROM tg_users WHERE chat_id = $1", chatID)
	err := row.Scan(&user.userName, &user.money, &user.latitude, &user.longitude)

	if err != nil {
		return nil, err
	}
	user.chatId = int64(chatID)
	return &user, nil
}

func fetchСallBackUser(ctx context.Context, pool *pgxpool.Pool, update tgbotapi.Update) (*User, error) {
	user := User{}
	chatID := update.CallbackQuery.Message.Chat.ID
	row := pool.QueryRow(ctx, "SELECT user_name, money,latitude,longitude FROM tg_users WHERE chat_id = $1", chatID)
	err := row.Scan(&user.userName, &user.money)
	if err != nil {
		return nil, err
	}
	user.chatId = chatID
	return &user, nil
}

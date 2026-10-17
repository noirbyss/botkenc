package main

import (
	"context"

	"github.com/noirbyss/botkenc/internal/transport/bot"
)

type Service struct{}

func (s Service) HasActiveSubscription(userID string) (bool, error) {
	return false, nil
}

func (s Service) UserIDByTelegramID(telegramID int64) (string, error) {
	return "1", nil
}

func (s Service) UserBalance(userID string) (int, error) {
	return 100, nil
}

func main() {
	bot, err := bot.New(context.Background(), Service{})
	if err != nil {
		panic(err)
	}

	if err := bot.Start(context.Background()); err != nil {
		panic(err)
	}
}

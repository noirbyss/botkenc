package main

import (
	"context"

	"github.com/noirbyss/botkenc/internal/transport/bot"
)

func main() {
	bot, err := bot.New(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	if err := bot.Start(context.Background()); err != nil {
		panic(err)
	}
}

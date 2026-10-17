package bot

import (
	"context"
	"fmt"
	"os"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

type Service interface {
	HasActiveSubscription(userID string) (bool, error)
	UserIDByTelegramID(telegramID int64) (string, error)
	UserBalance(userID string) (int, error) // TODO: поменять на float
}

type Bot struct {
	bot     *telego.Bot
	service Service
}

func New(ctx context.Context, service Service) (*Bot, error) {
	const envTelegramBotToken = "TELEGRAM_BOT_TOKEN"

	token := os.Getenv(envTelegramBotToken)
	if token == "" {
		return nil, fmt.Errorf("env %s not init", envTelegramBotToken)
	}

	bot, err := telego.NewBot(
		token,
		telego.WithDefaultDebugLogger(),
		telego.WithHealthCheck(ctx),
	)
	if err != nil {
		return nil, fmt.Errorf("create new bot: %w", err)
	}

	return &Bot{
		bot:     bot,
		service: service,
	}, nil
}

// TODO:
// - Измениить способ соединения
func (b *Bot) Start(ctx context.Context) error {
	updates, err := b.bot.UpdatesViaLongPolling(ctx, nil)
	if err != nil {
		return fmt.Errorf("connection setup: %w", err)
	}

	bh, err := th.NewBotHandler(b.bot, updates)
	if err != nil {
		return fmt.Errorf("create bot handler: %w", err)
	}

	b.registerHandlers(bh)

	defer func() {
		_ = bh.Stop()
	}()

	if err := bh.Start(); err != nil {
		return fmt.Errorf("run bot handler: %w", err)
	}

	return nil
}

func (b *Bot) registerHandlers(bh *th.BotHandler) {
	// TODO:
	// - Если b.StartHandler == nil - тогда bh.HandleMessage бросает панику.
	// Нужно добавить обработку.
	bh.HandleMessage(b.startHandler, th.CommandEqual("start"))
	bh.HandleCallbackQuery(b.mainHandler, th.CallbackDataEqual(string(mainMenu)))
}

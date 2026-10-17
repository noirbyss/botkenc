package bot

import (
	"context"
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

const (
	startMenu MenuID = "start"
	mainMenu  MenuID = "main"
)

type MenuID string

type Menu struct {
	id             MenuID
	text           string
	inlineKeyboard *telego.InlineKeyboardMarkup
	chatID         telego.ChatID
}

func (b *Bot) routeMenu(ctx context.Context, menu Menu) error {
	if menu.id == startMenu {
		_, err := b.bot.SendMessage(
			ctx,
			tu.Message(
				menu.chatID,
				menu.text,
			).WithReplyMarkup(menu.inlineKeyboard).WithParseMode(telego.ModeMarkdownV2),
		)

		if err != nil {
			return fmt.Errorf("send message: %w", err)
		}
	}

	return nil
}

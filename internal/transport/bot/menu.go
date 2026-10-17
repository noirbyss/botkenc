package bot

import (
	"context"
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

const (
	startMenu        MenuID = "start"
	mainMenu         MenuID = "main"
	subscriptionMenu MenuID = "subscription"
	promocodeMenu    MenuID = "promocode"
	affiliateMenu    MenuID = "affiliate"
	instructionsMenu MenuID = "instructions"
	supportMenu      MenuID = "support"
	tarrifsMenu      MenuID = "tarrifs"
)

type MenuID string

type Menu struct {
	id             MenuID
	text           string
	inlineKeyboard *telego.InlineKeyboardMarkup
	chatID         telego.ChatID
	messageID      *int
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
	} else {
		if menu.messageID == nil {
			return fmt.Errorf("message id no set")
		}

		messageID := *menu.messageID

		_, err := b.bot.EditMessageText(
			ctx,
			tu.EditMessageText(
				menu.chatID,
				messageID,
				menu.text,
			).WithReplyMarkup(menu.inlineKeyboard).WithParseMode(telego.ModeMarkdownV2),
		)
		if err != nil {
			return fmt.Errorf("edit message: %w", err)
		}
	}

	return nil
}

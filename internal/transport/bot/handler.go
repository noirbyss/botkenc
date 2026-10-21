package bot

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	"github.com/noirbyss/botkenc/internal/domain"
)

var (
	ErrGetMenu     = errors.New("get menu error")
	ErrDisplayMenu = errors.New("display menu")
)

func (b *Bot) startMessageHandler(ctx *th.Context, msg telego.Message) error {
	var (
		tgUID    = msg.From.ID
		screenID = domain.StartScreen
	)

	menu, err := b.backend.Menu(tgUID, screenID, nil)
	if err != nil {
		return fmt.Errorf("%w %w", ErrGetMenu, err)
	}

	if err := b.displayMenu(ctx.Context(), menu, tgUID, nil); err != nil {
		return fmt.Errorf("%w: %w", ErrDisplayMenu, err)
	}

	return nil
}

func (b *Bot) callbackHandler(ctx *th.Context, query telego.CallbackQuery) error {
	var (
		tgUID          = query.From.ID
		messageID      = query.Message.GetMessageID()
		screenID, args = parseCallbackData(query.Data)
	)

	menu, err := b.backend.Menu(tgUID, screenID, args)
	if err != nil {
		return fmt.Errorf("%w %w", ErrGetMenu, err)
	}

	if err := b.displayMenu(ctx.Context(), menu, tgUID, &messageID); err != nil {
		return fmt.Errorf("%w: %w", ErrDisplayMenu, err)
	}

	return nil
}

func parseCallbackData(callbackData string) (domain.ScreenID, *string) {
	if callbackData == string(domain.BackScreen) {
		return domain.BackScreen, nil
	}

	i := strings.IndexByte(callbackData, '(')

	if i == -1 {
		return domain.ScreenID(callbackData), nil
	}

	var (
		menu = callbackData[:i]
		args = callbackData[i+1 : len(callbackData)-1]
	)

	return domain.ScreenID(menu), &args
}

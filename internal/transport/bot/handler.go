package bot

import (
	"fmt"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

// В хендлере собираем нужные данные и направляем на роутинг.
func (b *Bot) startHandler(ctx *th.Context, msg telego.Message) error {
	fmt.Println("----Invoke handler-----")

	const message = `
	👋 Привет\! Это *\[Название\]* — VPN, который работает *даже при белых списках*\.
	
	Когда мобильный интернет ограничен и открываются только _«разрешённые»_ сайты, мы помогаем вернуть доступ *ко всему остальному*\.

	⚡ *Быстрое подключение* за пару минут
	📱 *iPhone, Android, Windows, macOS*
	🔒 *Без логов* и лишних данных

	Нажми *«Начать»*, чтобы продолжить 👇
	`

	inlineKeyboard := tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Начать").WithCallbackData(string(mainMenu)),
		),
	)

	menu := Menu{
		id:             startMenu,
		text:           message,
		inlineKeyboard: inlineKeyboard,
		chatID:         tu.ID(msg.From.ID),
	}

	if err := b.routeMenu(ctx.Context(), menu); err != nil {
		return fmt.Errorf("route menu: %w", err)
	}

	return nil
}

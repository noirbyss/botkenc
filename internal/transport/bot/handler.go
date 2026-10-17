package bot

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

var (
	ErrBackend   = errors.New("backend error")
	ErrRouteMenu = errors.New("route menu error")
)

// В хендлере собираем нужные данные и направляем на роутинг.
func (b *Bot) startHandler(ctx *th.Context, msg telego.Message) error {
	message := `
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
		return fmt.Errorf("%w: %w", ErrRouteMenu, err)
	}

	return nil
}

func (b *Bot) mainHandler(ctx *th.Context, query telego.CallbackQuery) error {
	telegramUserID := query.From.ID

	userID, err := b.service.UserIDByTelegramID(telegramUserID)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrBackend, err)
	}

	hasActiveSubscription, err := b.service.HasActiveSubscription(userID)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrBackend, err)
	}

	balance, err := b.service.UserBalance(userID)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrBackend, err)
	}

	var (
		message          string
		inlineKeyboard   *telego.InlineKeyboardMarkup
		mainRowsKeyboard = [][]telego.InlineKeyboardButton{
			tu.InlineKeyboardRow(
				tu.InlineKeyboardButton("Ввести промокод").WithCallbackData(string(promocodeMenu)),
			),
			tu.InlineKeyboardRow(
				tu.InlineKeyboardButton("Партнерская программа").WithCallbackData(string(affiliateMenu)),
			),
			tu.InlineKeyboardRow(
				tu.InlineKeyboardButton("Инструкции").WithCallbackData(string(instructionsMenu)),
			),
			tu.InlineKeyboardRow(
				tu.InlineKeyboardButton("Поддержка").WithCallbackData(string(supportMenu)),
			),
		}
	)
	if hasActiveSubscription {
		message = fmt.Sprintf(
			`
			👋 *С возвращением в \[Название\]\!*

			✅ Подписка *активна*, VPN готов к работе\.

			🆔 ID: %s
			💰 Баланс: %d ₽

			👇 Выберите действие:
			`,
			userID,
			balance,
		)

		rows := [][]telego.InlineKeyboardButton{
			tu.InlineKeyboardRow(
				tu.InlineKeyboardButton("Управление подпиской").WithCallbackData(string(subscriptionMenu)),
			),
		}

		rows = append(rows, mainRowsKeyboard...)
		inlineKeyboard = tu.InlineKeyboard(rows...)
	} else {
		message = fmt.Sprintf(
			`
👋 *Добро пожаловать в \[Название\]\!*

⚠️ У вас пока *нет активной подписки*\. Оформите её, чтобы подключиться к VPN и пользоваться интернетом без ограничений\.

🆔 ID: %s
💰 Баланс: %s ₽

👇 Выберите действие:
			`,
			userID,
			strconv.Itoa(balance),
		)

		rows := [][]telego.InlineKeyboardButton{
			tu.InlineKeyboardRow(
				tu.InlineKeyboardButton("Купиить подписку").WithCallbackData(string(tarrifsMenu)),
			),
		}

		rows = append(rows, mainRowsKeyboard...)
		inlineKeyboard = tu.InlineKeyboard(rows...)
	}

	messageID := query.Message.GetMessageID()

	menu := Menu{
		id:             mainMenu,
		text:           message,
		inlineKeyboard: inlineKeyboard,
		chatID:         tu.ID(telegramUserID),
		messageID:      &messageID,
	}

	if err := b.routeMenu(ctx.Context(), menu); err != nil {
		return fmt.Errorf("%w: %w", ErrRouteMenu, err)
	}

	return nil
}

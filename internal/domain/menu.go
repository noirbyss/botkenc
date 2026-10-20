package domain

import "github.com/mymmrac/telego"

const (
	StartMenu            MenuID = "start"
	MainMenuSubscribed   MenuID = "main_subscribed"
	MainMenuUnsubscribed MenuID = "main_unsubscribed"
	TariffsMenu          MenuID = "tariffs"
	SupportMenu          MenuID = "support"
	AffiliateMenu        MenuID = "affiliate"
	InstructionsMenu     MenuID = "Instructions"
	PaymentsMethodMenu   MenuID = "payments_method"
	SubscribeMenu        MenuID = "subscribe"
	DevicesMenu          MenuID = "devices"
	DevicesEmptyMenu     MenuID = "devices_empty"
	DevicesLimitMenu     MenuID = "devices_limit"
	BuyDeviceMenu        MenuID = "buy_device"
)

const (
	StartScreen          ScreenID = "start"
	MainScreen           ScreenID = "main"
	TariffsScreen        ScreenID = "tariffs"
	SupportScreen        ScreenID = "support"
	AffiliateScreen      ScreenID = "affiliate"
	InstructionsScreen   ScreenID = "Instructions"
	PaymentsMethodScreen ScreenID = "payments_method"
	SubscribeScreen      ScreenID = "subscribe"
	DevicesScreen        ScreenID = "devices"
	BuyDeviceScreen      ScreenID = "buy_device"
)

type MenuID string
type ScreenID string

type Menu struct {
	ID             MenuID
	TextMessage    string
	InlineKeyboard *telego.InlineKeyboardMarkup
}

func NewMenu(ID MenuID, text string, keyboard *telego.InlineKeyboardMarkup) *Menu {
	return &Menu{
		ID:             ID,
		TextMessage:    text,
		InlineKeyboard: keyboard,
	}
}

type MenuTemplate struct {
	ID                  MenuID
	TextMessageTemplate string
	InlineKeyboard      *telego.InlineKeyboardMarkup
}

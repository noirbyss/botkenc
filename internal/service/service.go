package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/noirbyss/botkenc/internal/domain"
)

const (
	placeholderUserID      string = "user_id"
	placeholderUserBalance string = "balance"

	placeholderReferralPercent string = "referral_percent"
	placeholderReferralInvited string = "referral_invited"
	placeholderReferralPaid    string = "referral_paid"
	placeholderReferralEarned  string = "referral_earned"

	placeholderProductName  string = "product_name"
	placeholderProductPrice string = "product_price"

	placeholderSubscribeExpires      = "subscribe_expires"
	placeholderSubscribeDevicesUsed  = "subscribe_devices_used"
	placeholderSubscribeDevicesLimit = "subscribe_device_limt"
	placeholderSubscribeURL          = "subscribe_url"

	placeholderDevicesUsed  = "devices_used"
	placeholderDevicesLimit = "devices_limit"
	placeholderDeviceName   = "device_name"
	placeholderDeviceDateAd = "device_date_ad"
)

type Repository interface {
	UserIDByTelegramUserID(ctx context.Context, telegramUserID int64) (string, error)
	UserBalance(ctx context.Context, userID string) (int, error) // TODO: тип баланса
	HasActiveSubscribe(ctx context.Context, userID string) (bool, error)

	ReferralPercent(ctx context.Context, userID string) (int, error)
	ReferralStats(ctx context.Context, userID string) (domain.ReferralStats, error)

	ProductPrice(ctx context.Context, productName string) (int, error) // TODO: тип прайса

	Subscribe(ctx context.Context, userID string) (domain.Subscribe, error)

	UserDevices(ctx context.Context, userID string) ([]domain.UserDevice, error)
	UserDevicesLimit(ctx context.Context, userID string) (int, error)

	Menu(ctx context.Context, menuID domain.MenuID) (domain.MenuTemplate, error)
}

type Service struct {
	repository Repository
}

func New(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (b Service) Menu(
	ctx context.Context,
	telegramUserID int64,
	screenID domain.ScreenID,
	args *string, // name
) (*domain.Menu, error) {
	switch screenID {
	case domain.StartScreen:
		return b.parseStartMenu(ctx)
	case domain.MainScreen:
		return b.parseMainMenu(ctx, telegramUserID)
	case domain.TariffsScreen:
		return b.parseTariffsMenu(ctx)
	case domain.SupportScreen:
		return b.parseSupportMenu(ctx, telegramUserID)
	case domain.AffiliateScreen:
		return b.parseAffiliateMenu(ctx, telegramUserID)
	case domain.InstructionsScreen:
		return b.parseInstructionsMenu(ctx)
	case domain.PaymentsMethodScreen:
		return b.parsePaymentsMethodMenu(ctx, telegramUserID, *args)
	case domain.SubscribeScreen:
		return b.parseSubscribeMenu(ctx, telegramUserID)
	case domain.DevicesScreen:
		return b.parseDevicesMenu(ctx, telegramUserID)
	case domain.BuyDeviceScreen:
		return b.parseBuyDeviceMenu(ctx, telegramUserID)
	default:
		return nil, fmt.Errorf("invalid screenID")
	}
}

// TODO: вынести в метод MenuTemplate
func parseTextTemplate(textTemplate string, placeholderValues map[string]string) string {
	pairs := make([]string, 0, len(placeholderValues)*2)

	for template, replace := range placeholderValues {
		pairs = append(pairs, "{+"+template+"}", replace)
	}

	return strings.NewReplacer(pairs...).Replace(textTemplate)
}

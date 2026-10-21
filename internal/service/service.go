package service

import (
	"context"
	"fmt"

	"github.com/noirbyss/botkenc/internal/domain"
)

const (
	placeholderUserID      string = "user_id"
	placeholderUserBalance string = "balance"

	placeholderReferralPercent   string = "referral_percent"
	placeholderReferralInvited   string = "referral_invited"
	placeholderReferralPaid      string = "referral_paid"
	placeholderReferralEarned    string = "referral_earned"
	placeholderReferralInviteURL string = "referral_invite_url"

	placeholderProductName  string = "product_name"
	placeholderProductPrice string = "product_price"

	placeholderSubscribeExpires      = "subscribe_expires"
	placeholderSubscribeDevicesUsed  = "subscribe_devices_used"
	placeholderSubscribeDevicesLimit = "subscribe_device_limt"
	placeholderSubscribeURL          = "subscribe_url"

	placeholderDeviceName   = "device_name"
	placeholderDeviceDateAd = "device_date_ad"
)

type Repository interface {
	UserIDByTelegramUserID(ctx context.Context, telegramUserID int64) (string, error)
	UserBalance(ctx context.Context, userID string) (int, error) // TODO: тип баланса
	HasActiveSubscribe(ctx context.Context, userID string) (bool, error)

	ReferralPercent(ctx context.Context, userID string) (int, error)
	ReferralStats(ctx context.Context, userID string) (domain.ReferralStats, error)
	ReferralInviteURL(ctx context.Context, userID string) (string, error)

	ProductPrice(ctx context.Context, productName string) (int, error) // TODO: тип прайса

	Subscribe(ctx context.Context, userID string) (domain.Subscribe, error)

	UserDevices(ctx context.Context, userID string) ([]domain.UserDevice, error)
	UserDevicesLimit(ctx context.Context, userID string) (int, error)

	Menu(ctx context.Context, menuID domain.MenuID) (domain.MenuTemplate, error)
}

type Service struct {
	repository Repository
	navigator  *navigator
}

func New(repository Repository) *Service {
	return &Service{
		repository: repository,
		navigator:  newNavigator(),
	}
}

func (s Service) Menu(
	ctx context.Context,
	telegramUserID int64,
	screenID domain.ScreenID,
	args *string, // name
) (*domain.Menu, error) {
	if screenID == domain.BackScreen {
		display := s.navigator.back(telegramUserID)
		return s.render(ctx, telegramUserID, display)
	}

	screen := domain.ScreenRef{
		ID: screenID,
	}

	menu, err := s.render(ctx, telegramUserID, screen)
	if err != nil {
		return nil, err
	}

	if screenID != domain.StartScreen {
		s.navigator.push(telegramUserID, screen)
	}

	return menu, nil
}

func (s *Service) render(
	ctx context.Context,
	telegramUserID int64,
	screen domain.ScreenRef,
) (*domain.Menu, error) {
	switch screen.ID {
	case domain.StartScreen:
		return s.renderStartScreen(ctx)
	case domain.MainScreen:
		return s.renderMainScreen(ctx, telegramUserID)
	case domain.TariffsScreen:
		return s.renderTariffsScreen(ctx)
	case domain.SupportScreen:
		return s.renderSupportScreen(ctx, telegramUserID)
	case domain.AffiliateScreen:
		return s.renderAffiliateScreen(ctx, telegramUserID)
	case domain.InstructionsScreen:
		return s.renderInstructionsScreen(ctx)
	case domain.PaymentsMethodScreen:
		if screen.Args == "" {
			return nil, fmt.Errorf("screen %s requires product name", screen.ID)
		}
		return s.renderPaymentsMethodScreen(ctx, telegramUserID, screen.Args)
	case domain.SubscribeScreen:
		return s.renderSubscribeScreen(ctx, telegramUserID)
	case domain.DevicesScreen:
		return s.renderDevicesScreen(ctx, telegramUserID)
	case domain.BuyDeviceScreen:
		return s.renderBuyDeviceScreen(ctx, telegramUserID)
	default:
		return nil, fmt.Errorf("unknown screen %q", screen.ID)
	}
}

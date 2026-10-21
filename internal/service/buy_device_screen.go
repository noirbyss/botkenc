package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/noirbyss/botkenc/internal/domain"
)

func (s Service) renderBuyDeviceScreen(ctx context.Context, telegramUserID int64) (*domain.Menu, error) {
	uid, err := s.repository.UserIDByTelegramUserID(ctx, telegramUserID)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGetUserID, err)
	}

	deviceLimit, err := s.repository.UserDevicesLimit(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGetUserDevicesLimit, err)
	}

	menuTemplate, err := s.repository.Menu(ctx, domain.BuyDeviceMenu)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGetMenuTemplate, err)
	}

	placeholderValues := make(map[string]string, 1)
	placeholderValues[placeholderSubscribeDevicesLimit] = strconv.Itoa(deviceLimit)

	menuTemplate.ParseTextTemplate(placeholderValues)

	return domain.NewMenu(
		menuTemplate.ID,
		menuTemplate.TextMessageTemplate,
		menuTemplate.InlineKeyboard,
	), nil
}

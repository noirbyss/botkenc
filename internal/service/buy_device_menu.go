package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/noirbyss/botkenc/internal/domain"
)

func (b Service) parseBuyDeviceMenu(ctx context.Context, telegramUserID int64) (*domain.Menu, error) {
	uid, err := b.repository.UserIDByTelegramUserID(ctx, telegramUserID)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGetUserID, err)
	}

	deviceLimit, err := b.repository.UserDevicesLimit(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGetUserDevicesLimit, err)
	}

	menuTemplate, err := b.repository.Menu(ctx, domain.BuyDeviceMenu)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGetMenuTemplate, err)
	}

	placeholderValues := make(map[string]string, 1)
	placeholderValues[placeholderDevicesLimit] = strconv.Itoa(deviceLimit)

	menuTemplate.TextMessageTemplate = parseTextTemplate(menuTemplate.TextMessageTemplate, placeholderValues)

	return domain.NewMenu(
		menuTemplate.ID,
		menuTemplate.TextMessageTemplate,
		menuTemplate.InlineKeyboard,
	), nil
}

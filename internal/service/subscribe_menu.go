package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/noirbyss/botkenc/internal/domain"
)

func (b Service) parseSubscribeMenu(ctx context.Context, telegramUserID int64) (*domain.Menu, error) {
	uid, err := b.repository.UserIDByTelegramUserID(ctx, telegramUserID)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGetUserID, err)
	}

	subscribeInfo, err := b.repository.Subscribe(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGetSubscribeInfo, err)
	}

	menuTemplate, err := b.repository.Menu(ctx, domain.SubscribeMenu)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGetMenuTemplate, err)
	}

	placeholderValues := make(map[string]string, 3)
	placeholderValues[placeholderSubscribeExpires] = subscribeInfo.ExpiresAt.Format("02.01.2006 03:04")
	placeholderValues[placeholderSubscribeDevicesUsed] = strconv.Itoa(subscribeInfo.DevicesUsed)
	placeholderValues[placeholderSubscribeDevicesLimit] = strconv.Itoa(subscribeInfo.DevicesLimit)
	placeholderValues[placeholderSubscribeURL] = subscribeInfo.URLConnection

	menuTemplate.TextMessageTemplate = parseTextTemplate(menuTemplate.TextMessageTemplate, placeholderValues)

	return domain.NewMenu(
		menuTemplate.ID,
		menuTemplate.TextMessageTemplate,
		menuTemplate.InlineKeyboard,
	), nil
}

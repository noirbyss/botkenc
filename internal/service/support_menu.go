package service

import (
	"context"
	"fmt"

	"github.com/noirbyss/botkenc/internal/domain"
)

func (b Service) parseSupportMenu(ctx context.Context, telegramUserID int64) (*domain.Menu, error) {
	uid, err := b.repository.UserIDByTelegramUserID(ctx, telegramUserID)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGetUserID, err)
	}

	menuTemplate, err := b.repository.Menu(ctx, domain.SupportMenu)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGetMenuTemplate, err)
	}

	placeholderValues := make(map[string]string, 1)
	placeholderValues[placeholderUserID] = uid

	menuTemplate.TextMessageTemplate = parseTextTemplate(menuTemplate.TextMessageTemplate, placeholderValues)

	return domain.NewMenu(
		menuTemplate.ID,
		menuTemplate.TextMessageTemplate,
		menuTemplate.InlineKeyboard,
	), nil
}

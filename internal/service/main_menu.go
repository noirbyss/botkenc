package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/noirbyss/botkenc/internal/domain"
)

func (b Service) parseMainMenu(ctx context.Context, telegramUserID int64) (*domain.Menu, error) {
	uid, err := b.repository.UserIDByTelegramUserID(ctx, telegramUserID)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGetUserID, err)
	}

	balance, err := b.repository.UserBalance(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGetUserBalance, err)
	}

	hasSubscribe, err := b.repository.HasActiveSubscribe(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrSubscriptionCheck, err)
	}

	var menuTemplate domain.MenuTemplate

	if hasSubscribe {
		menuTemplate, err = b.repository.Menu(ctx, domain.MainMenuSubscribed)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", ErrGetMenuTemplate, err)
		}
	} else {
		menuTemplate, err = b.repository.Menu(ctx, domain.MainMenuUnsubscribed)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", ErrGetMenuTemplate, err)
		}
	}

	placeholderValues := make(map[string]string, 2)
	placeholderValues[placeholderUserID] = uid
	placeholderValues[placeholderUserBalance] = strconv.Itoa(balance)

	menuTemplate.TextMessageTemplate = parseTextTemplate(menuTemplate.TextMessageTemplate, placeholderValues)

	return domain.NewMenu(
		menuTemplate.ID,
		menuTemplate.TextMessageTemplate,
		menuTemplate.InlineKeyboard,
	), nil
}

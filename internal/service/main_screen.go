package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/noirbyss/botkenc/internal/domain"
)

func (s Service) renderMainScreen(ctx context.Context, telegramUserID int64) (*domain.Menu, error) {
	uid, err := s.repository.UserIDByTelegramUserID(ctx, telegramUserID)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGetUserID, err)
	}

	balance, err := s.repository.UserBalance(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGetUserBalance, err)
	}

	hasSubscribe, err := s.repository.HasActiveSubscribe(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrSubscriptionCheck, err)
	}

	var menuTemplate domain.MenuTemplate

	if hasSubscribe {
		menuTemplate, err = s.repository.Menu(ctx, domain.MainMenuSubscribed)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", ErrGetMenuTemplate, err)
		}
	} else {
		menuTemplate, err = s.repository.Menu(ctx, domain.MainMenuUnsubscribed)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", ErrGetMenuTemplate, err)
		}
	}

	placeholderValues := make(map[string]string, 2)
	placeholderValues[placeholderUserID] = uid
	placeholderValues[placeholderUserBalance] = strconv.Itoa(balance)

	menuTemplate.ParseTextTemplate(placeholderValues)

	return domain.NewMenu(
		menuTemplate.ID,
		menuTemplate.TextMessageTemplate,
		menuTemplate.InlineKeyboard,
	), nil
}

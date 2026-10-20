package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/noirbyss/botkenc/internal/domain"
)

func (b Service) parsePaymentsMethodMenu(
	ctx context.Context,
	telegramUserID int64,
	productName string,
) (*domain.Menu, error) {
	uid, err := b.repository.UserIDByTelegramUserID(ctx, telegramUserID)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGetUserID, err)
	}

	balance, err := b.repository.UserBalance(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGetUserBalance, err)
	}

	productPrice, err := b.repository.ProductPrice(ctx, productName)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGetProductPrice, err)
	}

	templateMenu, err := b.repository.Menu(ctx, domain.PaymentsMethodMenu)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGetMenuTemplate, err)
	}

	placeholderValues := make(map[string]string, 3)
	placeholderValues[placeholderProductName] = productName
	placeholderValues[placeholderProductPrice] = strconv.Itoa(productPrice)
	placeholderValues[placeholderUserBalance] = strconv.Itoa(balance)

	templateMenu.TextMessageTemplate = parseTextTemplate(templateMenu.TextMessageTemplate, placeholderValues)

	return domain.NewMenu(
		templateMenu.ID,
		templateMenu.TextMessageTemplate,
		templateMenu.InlineKeyboard,
	), nil
}

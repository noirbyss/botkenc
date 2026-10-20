package service

import (
	"context"
	"fmt"

	"github.com/noirbyss/botkenc/internal/domain"
)

func (b Service) parseStartMenu(ctx context.Context) (*domain.Menu, error) {
	menuTemplate, err := b.repository.Menu(ctx, domain.StartMenu)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGetMenuTemplate, err)
	}

	return domain.NewMenu(
		menuTemplate.ID,
		menuTemplate.TextMessageTemplate,
		menuTemplate.InlineKeyboard,
	), nil
}

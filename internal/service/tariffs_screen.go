package service

import (
	"context"
	"fmt"

	"github.com/noirbyss/botkenc/internal/domain"
)

func (s Service) renderTariffsScreen(ctx context.Context) (*domain.Menu, error) {
	menuTemplate, err := s.repository.Menu(ctx, domain.TariffsMenu)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGetMenuTemplate, err)
	}

	return domain.NewMenu(
		menuTemplate.ID,
		menuTemplate.TextMessageTemplate,
		menuTemplate.InlineKeyboard,
	), nil
}

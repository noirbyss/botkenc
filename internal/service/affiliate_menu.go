package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/noirbyss/botkenc/internal/domain"
)

func (b Service) parseAffiliateMenu(ctx context.Context, telegramUserID int64) (*domain.Menu, error) {
	uid, err := b.repository.UserIDByTelegramUserID(ctx, telegramUserID)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGetUserID, err)
	}

	referralPercent, err := b.repository.ReferralPercent(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGetReferralPercent, err)
	}

	referralStats, err := b.repository.ReferralStats(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGetReferralStats, err)
	}

	menuTemplate, err := b.repository.Menu(ctx, domain.AffiliateMenu)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGetMenuTemplate, err)
	}

	placeholderValues := make(map[string]string, 4)
	placeholderValues[placeholderReferralPercent] = strconv.Itoa(referralPercent)
	placeholderValues[placeholderReferralInvited] = strconv.Itoa(referralStats.Invited)
	placeholderValues[placeholderReferralPaid] = strconv.Itoa(referralStats.Paid)
	placeholderValues[placeholderReferralEarned] = strconv.Itoa(referralStats.Earned)

	menuTemplate.TextMessageTemplate = parseTextTemplate(menuTemplate.TextMessageTemplate, placeholderValues)

	return domain.NewMenu(
		menuTemplate.ID,
		menuTemplate.TextMessageTemplate,
		menuTemplate.InlineKeyboard,
	), nil
}

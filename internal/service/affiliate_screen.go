package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/noirbyss/botkenc/internal/domain"
)

func (s Service) renderAffiliateScreen(ctx context.Context, telegramUserID int64) (*domain.Menu, error) {
	uid, err := s.repository.UserIDByTelegramUserID(ctx, telegramUserID)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGetUserID, err)
	}

	referralPercent, err := s.repository.ReferralPercent(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGetReferralPercent, err)
	}

	referralStats, err := s.repository.ReferralStats(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGetReferralStats, err)
	}

	referralInviteURL, err := s.repository.ReferralInviteURL(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGetReferralInviteURL, err)
	}

	menuTemplate, err := s.repository.Menu(ctx, domain.AffiliateMenu)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGetMenuTemplate, err)
	}

	placeholderValues := make(map[string]string, 4)
	placeholderValues[placeholderReferralPercent] = strconv.Itoa(referralPercent)
	placeholderValues[placeholderReferralInvited] = strconv.Itoa(referralStats.Invited)
	placeholderValues[placeholderReferralPaid] = strconv.Itoa(referralStats.Paid)
	placeholderValues[placeholderReferralEarned] = strconv.Itoa(referralStats.Earned)
	placeholderValues[placeholderReferralInviteURL] = referralInviteURL

	menuTemplate.ParseTextTemplate(placeholderValues)

	return domain.NewMenu(
		menuTemplate.ID,
		menuTemplate.TextMessageTemplate,
		menuTemplate.InlineKeyboard,
	), nil
}

package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/noirbyss/botkenc/internal/domain"
)

func (b Service) parseDevicesMenu(ctx context.Context, telegramUserID int64) (*domain.Menu, error) {
	uid, err := b.repository.UserIDByTelegramUserID(ctx, telegramUserID)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGetUserID, err)
	}

	userDevices, err := b.repository.UserDevices(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGetUserDevices, err)
	}

	devicesLimit, err := b.repository.UserDevicesLimit(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGetUserDevicesLimit, err)
	}

	var menuTemplate domain.MenuTemplate

	if len(userDevices) == 0 {
		menuTemplate, err = b.repository.Menu(ctx, domain.DevicesEmptyMenu)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", ErrGetMenuTemplate, err)
		}

		return domain.NewMenu(
			menuTemplate.ID,
			menuTemplate.TextMessageTemplate,
			menuTemplate.InlineKeyboard,
		), nil
	}

	if len(userDevices) >= devicesLimit {
		menuTemplate, err = b.repository.Menu(ctx, domain.DevicesLimitMenu)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", ErrGetMenuTemplate, err)
		}

		placeholderValues := make(map[string]string, 2)
		placeholderValues[placeholderDevicesUsed] = strconv.Itoa(len(userDevices))
		placeholderValues[placeholderDevicesLimit] = strconv.Itoa(devicesLimit)

		menuTemplate.TextMessageTemplate = parseTextTemplate(menuTemplate.TextMessageTemplate, placeholderValues)

		return domain.NewMenu(
			menuTemplate.ID,
			menuTemplate.TextMessageTemplate,
			menuTemplate.InlineKeyboard,
		), nil
	}

	placeholderValues := make(map[string]string, len(userDevices))

	for i, device := range userDevices {
		i++

		placeholderDeviceNameWithIndex := placeholderDeviceName + strconv.Itoa(i)
		placeholderDeviceDateAdWithIndex := placeholderDeviceDateAd + strconv.Itoa(i)

		placeholderValues[placeholderDeviceNameWithIndex] = device.Name
		placeholderValues[placeholderDeviceDateAdWithIndex] = device.DateAd.Format("02.01.2006")

		menuTemplate.TextMessageTemplate = parseTextTemplate(menuTemplate.TextMessageTemplate, placeholderValues)
	}

	placeholderValues[placeholderDevicesUsed] = strconv.Itoa(len(userDevices))
	placeholderValues[placeholderDevicesLimit] = strconv.Itoa(devicesLimit)

	menuTemplate.TextMessageTemplate = parseTextTemplate(menuTemplate.TextMessageTemplate, placeholderValues)

	return domain.NewMenu(
		menuTemplate.ID,
		menuTemplate.TextMessageTemplate,
		menuTemplate.InlineKeyboard,
	), nil
}

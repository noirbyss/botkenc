package service

import "errors"

var (
	ErrGetUserID           = errors.New("get user id error")
	ErrGetUserBalance      = errors.New("get user balance error")
	ErrGetMenuTemplate     = errors.New("get menu template error")
	ErrSubscriptionCheck   = errors.New("subscription check error")
	ErrGetReferralPercent  = errors.New("get referral percent error")
	ErrGetReferralStats    = errors.New("get referral stats error")
	ErrGetProductPrice     = errors.New("get prodict price error")
	ErrGetSubscribeInfo    = errors.New("get subscribe info error")
	ErrGetUserDevices      = errors.New("get user devices error")
	ErrGetUserDevicesLimit = errors.New("get user devices limit error")
)

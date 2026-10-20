package domain

import "time"

type Subscribe struct {
	ExpiresAt     time.Time
	DevicesUsed   int
	DevicesLimit  int
	URLConnection string
}

package notification

import "errors"

var (
	ErrNotificationAddressNotFound = errors.New("notification service address not found")
	ErrorNotificationService       = errors.New("failed to send notification")
)

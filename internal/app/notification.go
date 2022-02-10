package app

import "context"

type NotificationGRPCRepository interface {
	SendNotification(context.Context, User) (int64, error)
}

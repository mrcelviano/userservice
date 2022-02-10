package rpc

import (
	"context"
	"github.com/mrcelviano/userservice/commons"
	"github.com/mrcelviano/userservice/internal/app"
	"log"
)

const (
	ErrIn                   = "Error in"
	NotificationServiceName = "notificationservice"
)

type NotificationGRPC struct {
	pool commons.Pool
}

func NewNotificationGRPCRepository() app.NotificationGRPCRepository {
	defer log.Println("new notification service")
	pool, err := commons.InitGRPCConnectionPool(NotificationServiceName, 10)
	if err != nil {
		panic(err)
	}

	return &NotificationGRPC{pool: pool}
}

func (n *NotificationGRPC) SendNotification(ctx context.Context, user app.User) (taskID int64, err error) {
	return
}

func (n *NotificationGRPC) getClient() context.Context

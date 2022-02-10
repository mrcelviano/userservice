package rpc

import (
	"context"
	"github.com/mrcelviano/userservice/commons"
	"github.com/mrcelviano/userservice/internal/app"
	p "github.com/mrcelviano/userservice/proto"
	"github.com/pkg/errors"
	"log"
	"time"
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

func (n *NotificationGRPC) SendNotification(ctx context.Context, user app.User) (int64, error) {
	ctx, notificationClient, err := n.getClient(ctx)
	if err != nil {
		return 0, err
	}
	id, err := notificationClient.SendNotification(ctx, &p.SendNotificationRequest{User: &p.User{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}})
	if err != nil {
		return 0, errors.Wrap(err, ErrIn+NotificationServiceName)
	}

	return id.TaskID, nil
}

func (n *NotificationGRPC) getClient(ctx context.Context) (context.Context, p.NotificationServiceClient, error) {
	ctx, _ = context.WithTimeout(ctx, time.Second*10)
	clientConn, err := n.pool.Get(ctx)
	if err != nil {
		return nil, nil, errors.Wrap(err, "can`t get connections")
	}
	return ctx, p.NewNotificationServiceClient(clientConn), nil
}

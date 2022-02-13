package notification

import (
	"context"
	client "github.com/mrcelviano/userservice/pkg/notification/proto"
	pool "github.com/processout/grpc-go-pool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

const (
	notificationServiceName = "notificationservice"
)

type notificationGRPC struct {
	pool *pool.Pool
}

func NewNotificationClient(serviceAddress map[string]string) (Service, error) {
	notification := notificationGRPC{}
	notificationAddress, err := notification.getServiceName(serviceAddress)
	if err != nil {
		return nil, err
	}
	err = notification.newGRPCConnectionPool(notificationAddress, 10)
	if err != nil {
		return nil, err
	}

	return &notification, nil
}

func (n *notificationGRPC) RegisterNotification(ctx context.Context, userID int64) (bool, error) {
	ctx, cansel := context.WithTimeout(ctx, time.Second*10)
	defer cansel()

	clientConn, err := n.pool.Get(ctx)
	if err != nil {
		return false, err
	}

	isRegistered, err := client.NewNotificationServiceClient(clientConn).RegisterNotification(ctx,
		&client.RegisterNotificationRequest{
			UserID: userID,
		})
	if err != nil {
		return false, ErrorNotificationService
	}

	return isRegistered.GetIsRegistered(), nil
}

func (n *notificationGRPC) newGRPCConnectionPool(serviceAddress string, capacity int) (err error) {
	factory := func() (*grpc.ClientConn, error) {
		conn, err := grpc.Dial(serviceAddress,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithDefaultCallOptions(
				grpc.WaitForReady(true),
			),
		)
		if err != nil {
			return nil, err
		}
		return conn, err
	}
	n.pool, err = pool.New(factory, 0, capacity*10, time.Minute, time.Minute)
	if err != nil {
		return err
	}
	return nil
}

func (n *notificationGRPC) getServiceName(serviceAddress map[string]string) (string, error) {
	notificationAddress, ok := serviceAddress[notificationServiceName]
	if !ok {
		return "", ErrNotificationAddressNotFound
	}
	return notificationAddress, nil
}

package notification

import (
	"context"
	"github.com/mrcelviano/userservice/pkg/notification/proto"
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

func (n *notificationGRPC) SendNotification(ctx context.Context, id int64, email, name string) (int64, error) {
	ctx, cansel := context.WithTimeout(ctx, time.Second*10)
	defer cansel()

	clientConn, err := n.pool.Get(ctx)
	if err != nil {
		return 0, err
	}

	taskID, err := proto.NewNotificationServiceClient(clientConn).SendNotification(ctx,
		&proto.SendNotificationRequest{
			User: &proto.User{
				ID:    id,
				Email: email,
				Name:  name,
			}})
	if err != nil {
		return 0, ErrorNotificationService
	}

	return taskID.GetTaskID(), nil
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

package notification

import (
	"context"
)

type Service interface {
	SendNotification(ctx context.Context, id int64, email, name string) (int64, error)
}

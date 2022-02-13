package notification

import (
	"context"
)

type Service interface {
	RegisterNotification(ctx context.Context, userID int64) (bool, error)
}

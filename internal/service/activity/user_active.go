package activity

import "context"

// UserActiveActivityRepo interface
type UserActiveActivityRepo interface {
	UserActive(ctx context.Context, userID string) (err error)
}

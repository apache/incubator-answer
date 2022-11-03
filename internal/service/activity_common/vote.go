package activity_common

import (
	"context"
)

// VoteRepo activity repository
type VoteRepo interface {
	GetVoteStatus(ctx context.Context, objectId, userId string) (status string)
	GetVoteCount(ctx context.Context, activityTypes []int) (count int64, err error)
}

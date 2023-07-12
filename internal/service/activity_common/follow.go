package activity_common

import "context"

type FollowRepo interface {
	GetFollowIDs(ctx context.Context, userID, objectType string) (followIDs []string, err error)
	GetFollowAmount(ctx context.Context, objectID string) (followAmount int, err error)
	GetFollowUserIDs(ctx context.Context, objectID string) (userIDs []string, err error)
	IsFollowed(ctx context.Context, userId, objectId string) (bool, error)
}

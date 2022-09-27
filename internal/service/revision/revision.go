package revision

import (
	"context"
	"github.com/segmentfault/answer/internal/entity"
	"xorm.io/xorm"
)

// RevisionRepo revision repository
type RevisionRepo interface {
	AddRevision(ctx context.Context, revision *entity.Revision, autoUpdateRevisionID bool) (err error)
	GetRevision(ctx context.Context, id string) (revision *entity.Revision, exist bool, err error)
	GetLastRevisionByObjectID(ctx context.Context, objectID string) (revision *entity.Revision, exist bool, err error)
	GetRevisionList(ctx context.Context, revision *entity.Revision) (revisionList []entity.Revision, err error)
	UpdateObjectRevisionId(ctx context.Context, revision *entity.Revision, session *xorm.Session) (err error)
}

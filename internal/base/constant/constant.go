package constant

import "time"

const (
	Default_PageSize           = 20        //Default number of pages
	Key_UserID                 = "_UserID" //session userid
	LoginUserID                = "login_user_id"
	LoginUserVerify            = "login_user_verify"
	UserStatusChangedCacheKey  = "answer:user:status:"
	UserStatusChangedCacheTime = 7 * 24 * time.Hour
	UserTokenCacheKey          = "answer:user:token:"
	UserTokenCacheTime         = 7 * 24 * time.Hour
	AdminTokenCacheKey         = "answer:admin:token:"
	AdminTokenCacheTime        = 7 * 24 * time.Hour
)

const (
	QuestionObjectType   = "question"
	AnswerObjectType     = "answer"
	TagObjectType        = "tag"
	UserObjectType       = "user"
	CollectionObjectType = "collection"
	CommentObjectType    = "comment"
	ReportObjectType     = "report"
)

// ObjectTypeStrMapping key => value
// object TagID AnswerList
// key equal database's table name
var (
	ObjectTypeStrMapping = map[string]int{
		QuestionObjectType:   1,
		AnswerObjectType:     2,
		TagObjectType:        3,
		UserObjectType:       4,
		CollectionObjectType: 6,
		CommentObjectType:    7,
		ReportObjectType:     8,
	}

	ObjectTypeNumberMapping = map[int]string{
		1: QuestionObjectType,
		2: AnswerObjectType,
		3: TagObjectType,
		4: UserObjectType,
		6: CollectionObjectType,
		7: CommentObjectType,
		8: ReportObjectType,
	}
)

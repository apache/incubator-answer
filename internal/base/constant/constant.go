package constant

import "time"

const (
	DefaultPageSize            = 20 // Default number of pages
	UserStatusChangedCacheKey  = "answer:user:status:"
	UserStatusChangedCacheTime = 7 * 24 * time.Hour
	UserTokenCacheKey          = "answer:user:token:"
	UserTokenCacheTime         = 7 * 24 * time.Hour
	AdminTokenCacheKey         = "answer:admin:token:"
	AdminTokenCacheTime        = 7 * 24 * time.Hour
	AcceptLanguageFlag         = "Accept-Language"
	UserTokenMappingCacheKey   = "answer:user-token:mapping:"
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
	Version string = ""

	PathIgnoreMap map[string]bool

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

const (
	SiteTypeGeneral   = "general"
	SiteTypeInterface = "interface"
	SiteTypeBranding  = "branding"
	SiteTypeWrite     = "write"
	SiteTypeLegal     = "legal"
	SiteTypeSeo       = "seo"
)

func ExistInPathIgnore(name string) bool {
	_, ok := PathIgnoreMap[name]
	if ok {
		return true
	}
	return false
}

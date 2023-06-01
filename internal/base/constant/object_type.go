package constant

const (
	QuestionObjectType   = "question"
	AnswerObjectType     = "answer"
	TagObjectType        = "tag"
	UserObjectType       = "user"
	CollectionObjectType = "collection"
	CommentObjectType    = "comment"
	ReportObjectType     = "report"
)

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

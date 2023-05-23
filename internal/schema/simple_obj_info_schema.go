package schema

// SimpleObjectInfo simple object info
type SimpleObjectInfo struct {
	ObjectID            string `json:"object_id"`
	ObjectCreatorUserID string `json:"object_creator_user_id"`
	QuestionID          string `json:"question_id"`
	QuestionStatus      int    `json:"status"`
	AnswerID            string `json:"answer_id"`
	CommentID           string `json:"comment_id"`
	TagID               string `json:"tag_id"`
	ObjectType          string `json:"object_type"`
	Title               string `json:"title"`
	Content             string `json:"content"`
}

type UnreviewedRevisionInfoInfo struct {
	ObjectID string     `json:"object_id"`
	Title    string     `json:"title"`
	Content  string     `json:"content"`
	Html     string     `json:"html"`
	Tags     []*TagResp `json:"tags"`
}

package schema

// SimpleObjectInfo simple object info
type SimpleObjectInfo struct {
	ObjectID      string `json:"object_id"`
	ObjectCreator string `json:"object_creator"`
	QuestionID    string `json:"question_id"`
	AnswerID      string `json:"answer_id"`
	CommentID     string `json:"comment_id"`
	TagID         string `json:"tag_id"`
	ObjectType    string `json:"object_type"`
	Title         string `json:"title"`
	Content       string `json:"content"`
}

package constant

// | 问题 回答 标签 | undeleted | 操作者        |              | 恢复删除的内容                                               |
// | 问题 回答 标签 | deleted   | 操作者        |              | 删除内容                                                     |
// | 问题 回答 标签 | rollback | 编辑者        | 显示编辑理由 | 回滚版本编辑记录； 点击 Type 显示最近的版本比较              |
// | 问题 回答 标签 | edit     | 编辑者        | 显示编辑理由 | 编辑记录； 点击 Type 显示最近的版本比较                      |
// | 问题 回答     | downvote  | 投票者 or N/A |              | 内容点踩，名字仅管理员可见； 取消时显示已取消和取消时间      |
// | 问题 回答     | upvote    | 投票者        |              | 内容点赞； 取消时显示已取消和取消时间                        |
// | 问题 回答     | accept    | 提问者        |              | 采纳答案，Type 链接到对应的回答； 取消时显示已取消和取消时间 |
// | 问题 回答     | commented | 评论者        | 显示评论内容 | 添加评论，Type 链接到对应的评论                              |
// | 问题         | answered  | 回答者        |              | 添加回答，Type 链接到对应的回答                              |
// | 问题         | reopened  | 操作者        |              | 重新开启问题                                                 |
// | 问题         | closed    | 操作者        | 显示关闭理由 | 关闭问题                                                     |
// | 问题         | asked    | 提问者        |              | 初始提问版本，点击展开无需比较                               |
// | 回答         | answered | 回答者        |              | 初始回答版本，点击展开无需比较                               |
// | 标签         | created  | 创建者        |              | 初始标签版本，点击展开无需比较                               |

// question activity

type ActivityTypeKey string

const (
	ActQuestionAsked     ActivityTypeKey = "question.asked"
	ActQuestionClosed    ActivityTypeKey = "question.closed"
	ActQuestionReopened  ActivityTypeKey = "question.reopened"
	ActQuestionAnswered  ActivityTypeKey = "question.answered"
	ActQuestionCommented ActivityTypeKey = "question.commented"
	ActQuestionAccept    ActivityTypeKey = "question.accept"
	ActQuestionUpvote    ActivityTypeKey = "question.upvote"
	ActQuestionDownvote  ActivityTypeKey = "question.downvote"
	ActQuestionEdit      ActivityTypeKey = "question.edit"
	ActQuestionRollback  ActivityTypeKey = "question.rollback"
	ActQuestionDeleted   ActivityTypeKey = "question.deleted"
	ActQuestionUndeleted ActivityTypeKey = "question.undeleted"
)

// answer activity

const (
	ActAnswerAnswered  ActivityTypeKey = "answer.answered"
	ActAnswerCommented ActivityTypeKey = "answer.commented"
	ActAnswerAccept    ActivityTypeKey = "answer.accept"
	ActAnswerUpvote    ActivityTypeKey = "answer.upvote"
	ActAnswerDownvote  ActivityTypeKey = "answer.downvote"
	ActAnswerEdit      ActivityTypeKey = "answer.edit"
	ActAnswerRollback  ActivityTypeKey = "answer.rollback"
	ActAnswerDeleted   ActivityTypeKey = "answer.deleted"
	ActAnswerUndeleted ActivityTypeKey = "answer.undeleted"
)

// tag activity

const (
	ActTagCreated   ActivityTypeKey = "tag.created"
	ActTagEdit      ActivityTypeKey = "tag.edit"
	ActTagRollback  ActivityTypeKey = "tag.rollback"
	ActTagDeleted   ActivityTypeKey = "tag.deleted"
	ActTagUndeleted ActivityTypeKey = "tag.undeleted"
)

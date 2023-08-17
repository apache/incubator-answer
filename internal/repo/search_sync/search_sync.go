package search_sync

import (
	"context"
	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/plugin"
)

func NewPluginSyncer(data *data.Data) plugin.SearchSyncer {
	return &PluginSyncer{data: data}
}

type PluginSyncer struct {
	data *data.Data
}

func (p *PluginSyncer) GetAnswersPage(ctx context.Context, page, pageSize int) (answerList []*entity.Answer, err error) {
	answerList = make([]*entity.Answer, 0)
	startNum := (page - 1) * pageSize
	err = p.data.DB.Context(ctx).Limit(pageSize, startNum).Find(&answerList)
	return answerList, err
}

func (p *PluginSyncer) GetQuestionsPage(ctx context.Context, page, pageSize int) (questionList []*entity.Question, err error) {
	questionList = make([]*entity.Question, 0)
	startNum := (page - 1) * pageSize
	err = p.data.DB.Context(ctx).Limit(pageSize, startNum).Find(&questionList)
	return questionList, err
}

package search_sync

import (
	"context"
	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/base/pager"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/plugin"
)

func NewPluginSyncer(data *data.Data) plugin.SearchSyncer {
	return &PluginSyncer{data: data}
}

type PluginSyncer struct {
	data *data.Data
}

func (p *PluginSyncer) GetAnswersPage(ctx context.Context, page, pageSize int) (answerList []*entity.Answer, total int64, err error) {
	answerList = make([]*entity.Answer, 0)
	total, err = pager.Help(page, pageSize, answerList, &entity.Answer{}, p.data.DB.Context(ctx))
	return answerList, total, err
}

func (p *PluginSyncer) GetQuestionsPage(ctx context.Context, page, pageSize int) (questionList []*entity.Question, total int64, err error) {
	questionList = make([]*entity.Question, 0)
	total, err = pager.Help(page, pageSize, questionList, &entity.Question{}, p.data.DB.Context(ctx))
	return questionList, total, err
}

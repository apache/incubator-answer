package schema

type SiteMapList struct {
	QuestionIDs []*SiteMapQuestionInfo `json:"question_ids"`
	MaxPageNum  []int                  `json:"max_page_num"`
}

type SiteMapPageList struct {
	PageData []*SiteMapQuestionInfo `json:"page_data"`
}

type SiteMapQuestionInfo struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	UpdateTime string `json:"time"`
}

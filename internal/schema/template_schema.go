package schema

import "time"

type Paginator struct {
	Pages      []int
	Totalpages int
	Prevpage   int
	Nextpage   int
	Currpage   int
}

type QAPageJsonLD struct {
	Context    string `json:"@context"`
	Type       string `json:"@type"`
	MainEntity struct {
		Type        string    `json:"@type"`
		Name        string    `json:"name"`
		Text        string    `json:"text"`
		AnswerCount int       `json:"answerCount"`
		UpvoteCount int       `json:"upvoteCount"`
		DateCreated time.Time `json:"dateCreated"`
		Author      struct {
			Type string `json:"@type"`
			Name string `json:"name"`
		} `json:"author"`
		AcceptedAnswer  AcceptedAnswerItem     `json:"acceptedAnswer"`
		SuggestedAnswer []*SuggestedAnswerItem `json:"suggestedAnswer"`
	} `json:"mainEntity"`
}

type AcceptedAnswerItem struct {
	Type        string `json:"@type"`
	Text        string `json:"text"`
	UpvoteCount int    `json:"upvoteCount"`
	URL         string `json:"url"`
	Author      struct {
		Type string `json:"@type"`
		Name string `json:"name"`
	} `json:"author"`
}

type SuggestedAnswerItem struct {
	Type        string    `json:"@type"`
	Text        string    `json:"text"`
	DateCreated time.Time `json:"dateCreated"`
	UpvoteCount int       `json:"upvoteCount"`
	URL         string    `json:"url"`
	Author      struct {
		Type string `json:"@type"`
		Name string `json:"name"`
	} `json:"author"`
}

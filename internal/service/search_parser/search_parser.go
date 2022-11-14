package search_parser

import (
	"context"
	"github.com/answerdev/answer/internal/schema"
	tagcommon "github.com/answerdev/answer/internal/service/tag_common"
	usercommon "github.com/answerdev/answer/internal/service/user_common"
	"github.com/answerdev/answer/pkg/converter"
	"regexp"
	"strings"
)

type SearchParser struct {
	tagRepo    tagcommon.TagRepo
	userCommon *usercommon.UserCommon
}

func NewSearchParser(tagRepo tagcommon.TagRepo, userCommon *usercommon.UserCommon) *SearchParser {
	return &SearchParser{
		tagRepo:    tagRepo,
		userCommon: userCommon,
	}
}

// ParseStructure parse search structure, maybe match one of type all/questions/answers,
// but if match two type, it will return false
func (sp *SearchParser) ParseStructure(dto *schema.SearchDTO) (
	searchType string,
	// search all
	userID string,
	votes int,
	// search questions
	notAccepted bool,
	isQuestion bool,
	views,
	answers int,
	// search answers
	accepted bool,
	questionID string,
	isAnswer bool,
	// common fields
	tags,
	words []string,
) {
	var (
		query         = dto.Query
		currentUserID = dto.UserID
		all           = 0
		q             = 0
		a             = 0
		withWords     = []string{}
		limitWords    = 5
	)

	// match tags
	tags = sp.parseTags(&query)

	// match all
	userID = sp.parseUserID(&query, currentUserID)
	if userID != "" {
		searchType = "all"
		all = 1
	}
	votes = sp.parseVotes(&query)
	if votes != -1 {
		searchType = "all"
		all = 1
	}
	withWords = sp.parseWithin(&query)
	if len(withWords) > 0 {
		searchType = "all"
		all = 1
	}

	// match questions
	notAccepted = sp.parseNotAccepted(&query)
	if notAccepted {
		searchType = "question"
		q = 1
	}
	isQuestion = sp.parseIsQuestion(&query)
	if isQuestion {
		searchType = "question"
		q = 1
	}
	views = sp.parseViews(&query)
	if views != -1 {
		searchType = "question"
		q = 1
	}
	answers = sp.parseAnswers(&query)
	if answers != -1 {
		searchType = "question"
		q = 1
	}

	// match answers
	accepted = sp.parseAccepted(&query)
	if accepted {
		searchType = "answer"
		a = 1
	}
	questionID = sp.parseQuestionID(&query)
	if questionID != "" {
		searchType = "answer"
		a = 1
	}
	isAnswer = sp.parseIsAnswer(&query)
	if isAnswer {
		searchType = "answer"
		a = 1
	}

	words = strings.Split(query, " ")
	if len(withWords) > 0 {
		words = append(withWords, words...)
	}

	// check limit words
	if len(words) > limitWords {
		words = words[:limitWords]
	}

	// check tags' search is all or question
	if len(tags) > 0 {
		if len(words) > 0 {
			searchType = "all"
			all = 1
		} else {
			searchType = "question"
			q = 1
		}
	}

	// check match types greater than 1
	if all+q+a > 1 {
		searchType = ""
	}

	// check not match
	if all+q+a == 0 && len(words) > 0 {
		searchType = "all"
	}

	return
}

// parseTags parse search tags, return tag ids array
func (sp *SearchParser) parseTags(query *string) (tags []string) {
	var (
		// expire tag pattern
		exprTag = `(?m)\[([a-zA-Z0-9-\+\.#]+)\]{1}?`
		q       = *query
		limit   = 5
	)

	re := regexp.MustCompile(exprTag)
	res := re.FindAllStringSubmatch(q, -1)
	if len(res) == 0 {
		return
	}
	tags = make([]string, len(res))
	for i, item := range res {
		tag, exists, err := sp.tagRepo.GetTagBySlugName(context.TODO(), item[1])
		if err != nil || !exists {
			continue
		}
		tags[i] = tag.ID
	}

	// limit maximum 5 tags
	if len(tags) > limit {
		tags = tags[:limit]
	}

	q = strings.TrimSpace(re.ReplaceAllString(q, ""))
	*query = q
	return
}

// parseUserID return user id or current login user id
func (sp *SearchParser) parseUserID(query *string, currentUserID string) (userID string) {
	var (
		exprUserID = `(?m)^user:([a-z0-9._-]+)`
		exprMe     = "user:me"
		q          = *query
	)

	re := regexp.MustCompile(exprUserID)
	res := re.FindStringSubmatch(q)
	if len(res) == 2 {
		name := res[1]
		user, has, err := sp.userCommon.GetUserBasicInfoByUserName(nil, name)
		if err == nil && has {
			userID = user.ID
			q = re.ReplaceAllString(q, "")
		}
	} else if strings.Index(q, exprMe) != -1 {
		userID = currentUserID
		q = strings.ReplaceAll(q, exprMe, "")
	}
	*query = strings.TrimSpace(q)
	return
}

// parseVotes return the votes of search query
func (sp *SearchParser) parseVotes(query *string) (votes int) {
	var (
		expr = `(?m)^score:([0-9]+)`
		q    = *query
	)
	votes = -1

	re := regexp.MustCompile(expr)
	res := re.FindStringSubmatch(q)
	if len(res) == 2 {
		votes = converter.StringToInt(res[1])
		q = re.ReplaceAllString(q, "")
	}

	*query = strings.TrimSpace(q)
	return
}

// parseWithin parse quotes within words like: "hello world"
func (sp *SearchParser) parseWithin(query *string) (words []string) {
	var (
		q    = *query
		expr = `(?U)(".+")`
	)
	re := regexp.MustCompile(expr)
	matches := re.FindAllStringSubmatch(q, -1)
	words = []string{}
	for _, match := range matches {
		words = append(words, match[1])
	}
	q = re.ReplaceAllString(q, "")
	*query = strings.TrimSpace(q)
	return
}

// parseNotAccepted return the question has not accepted the answer
func (sp *SearchParser) parseNotAccepted(query *string) (notAccepted bool) {
	var (
		q    = *query
		expr = `hasaccepted:no`
	)

	if strings.Index(q, expr) != -1 {
		q = strings.ReplaceAll(q, expr, "")
		notAccepted = true
	}

	*query = strings.TrimSpace(q)
	return
}

// parseIsQuestion check the result if only limit question or not
func (sp *SearchParser) parseIsQuestion(query *string) (isQuestion bool) {
	var (
		q    = *query
		expr = `is:question`
	)

	if strings.Index(q, expr) == 0 {
		q = strings.ReplaceAll(q, expr, "")
		isQuestion = true
	}

	*query = strings.TrimSpace(q)
	return
}

// parseViews check search has views or not
func (sp *SearchParser) parseViews(query *string) (views int) {
	var (
		q    = *query
		expr = `(?m)^views:([0-9]+)`
	)
	views = -1

	re := regexp.MustCompile(expr)
	res := re.FindStringSubmatch(q)
	if len(res) == 2 {
		views = converter.StringToInt(res[1])
		q = re.ReplaceAllString(q, "")
	}
	*query = strings.TrimSpace(q)
	return
}

// parseAnswers check whether specified answer count for question
func (sp *SearchParser) parseAnswers(query *string) (answers int) {
	var (
		q    = *query
		expr = `(?m)^answers:([0-9]+)`
	)
	answers = -1

	re := regexp.MustCompile(expr)
	res := re.FindStringSubmatch(q)
	if len(res) == 2 {
		answers = converter.StringToInt(res[1])
		q = re.ReplaceAllString(q, "")
	}

	*query = strings.TrimSpace(q)
	return
}

// parseAccepted check the search is limit accepted answer or not
func (sp *SearchParser) parseAccepted(query *string) (accepted bool) {
	var (
		q    = *query
		expr = `isaccepted:yes`
	)

	if strings.Index(q, expr) != -1 {
		accepted = true
		strings.ReplaceAll(q, expr, "")
	}

	*query = strings.TrimSpace(q)
	return
}

// parseQuestionID check whether specified question's id
func (sp *SearchParser) parseQuestionID(query *string) (questionID string) {
	var (
		q    = *query
		expr = `(?m)^inquestion:([0-9]+)`
	)

	re := regexp.MustCompile(expr)
	res := re.FindStringSubmatch(q)
	if len(res) == 2 {
		questionID = res[1]
		q = re.ReplaceAllString(q, "")
	}

	*query = strings.TrimSpace(q)
	return
}

// parseIsAnswer check the result if only limit answer or not
func (sp *SearchParser) parseIsAnswer(query *string) (isAnswer bool) {
	var (
		q    = *query
		expr = `is:answer`
	)

	if strings.Index(q, expr) != -1 {
		isAnswer = true
		q = strings.ReplaceAll(q, expr, "")
	}

	*query = strings.TrimSpace(q)
	return
}

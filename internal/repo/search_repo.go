package repo

import (
	"context"
	"strings"
	"time"

	"github.com/jinzhu/copier"
	"github.com/segmentfault/answer/internal/base/data"
	"github.com/segmentfault/answer/internal/base/reason"
	"github.com/segmentfault/answer/internal/entity"
	"github.com/segmentfault/answer/internal/schema"
	"github.com/segmentfault/answer/internal/service/search_common"
	"github.com/segmentfault/answer/internal/service/unique"
	usercommon "github.com/segmentfault/answer/internal/service/user_common"
	"github.com/segmentfault/answer/pkg/converter"
	"github.com/segmentfault/answer/pkg/obj"
	"github.com/segmentfault/pacman/errors"
	"xorm.io/builder"
)

// searchRepo tag repository
type searchRepo struct {
	data         *data.Data
	userCommon   *usercommon.UserCommon
	uniqueIDRepo unique.UniqueIDRepo
}

// NewSearchRepo new repository
func NewSearchRepo(data *data.Data, uniqueIDRepo unique.UniqueIDRepo, userCommon *usercommon.UserCommon) search_common.SearchRepo {
	return &searchRepo{
		data:         data,
		uniqueIDRepo: uniqueIDRepo,
		userCommon:   userCommon,
	}
}

func (sr *searchRepo) SearchContents(ctx context.Context, words []string, tagID, userID string, votes int, page, size int, order string) (resp []schema.SearchResp, total int64, err error) {
	var (
		b  *builder.Builder
		ub *builder.Builder
	)
	b = builder.MySQL().Select(
		"`question`.`id`",
		"`question`.`id` as `question_id`",
		"`title`",
		"`original_text`",
		"`question`.`created_at`",
		"`user_id`",
		"`vote_count`",
		"`answer_count`",
		"0 as `accepted`",
		"`question`.`status` as `status`",
		"`post_update_time`",
	).From("`question`")
	ub = builder.MySQL().Select(
		"`answer`.`id` as `id`",
		"`question_id`",
		"`question`.`title` as `title`",
		"`answer`.`original_text` as `original_text`",
		"`answer`.`created_at`",
		"`answer`.`user_id` as `user_id`",
		"`answer`.`vote_count` as `vote_count`",
		"0 as `answer_count`",
		"`adopted` as `accepted`",
		"`answer`.`status` as `status`",
		"`answer`.`created_at` as `post_update_time`",
	).From("`answer`").
		LeftJoin("`question`", "`question`.id = `answer`.question_id")

	for i, word := range words {
		if i == 0 {
			b.Where(builder.Like{"original_text", word})
			ub.Where(builder.Like{"`answer`.original_text", word})
		} else {
			b.Or(builder.Like{"original_text", word})
			ub.Or(builder.Like{"`answer`.original_text", word})
		}
	}

	// check tag
	if tagID != "" {
		b.Join("INNER", "tag_rel", "question.id = tag_rel.object_id").
			Where(builder.Eq{"tag_rel.tag_id": tagID})
	}

	// check user
	if userID != "" {
		b.Where(builder.Eq{"question.user_id": userID})
		ub.Where(builder.Eq{"answer.user_id": userID})
	}

	// check vote
	if votes == 0 {
		b.Where(builder.Eq{"question.vote_count": votes})
		ub.Where(builder.Eq{"answer.vote_count": votes})
	} else if votes > 0 {
		b.Where(builder.Gte{"question.vote_count": votes})
		ub.Where(builder.Gte{"answer.vote_count": votes})
	}

	b = b.Union("all", ub)
	_, _, err = b.ToSQL()
	if err != nil {
		return
	}

	res, err := sr.data.DB.Query(builder.MySQL().Select("*").From(b, "t").OrderBy(sr.parseOrder(ctx, order)).Limit(size, page-1))

	tr, err := sr.data.DB.Query(builder.MySQL().Select("count(*) total").From(b, "c"))
	if len(tr) != 0 {
		total = converter.StringToInt64(string(tr[0]["total"]))
	}
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
		return
	} else {
		resp, err = sr.parseResult(ctx, res)
		return
	}
}

func (sr *searchRepo) SearchQuestions(ctx context.Context, words []string, limitNoAccepted bool, answers, page, size int, order string) (resp []schema.SearchResp, total int64, err error) {
	b := builder.MySQL().Select(
		"`id`",
		"`id` as `question_id`",
		"`title`",
		"`original_text`",
		"`created_at`",
		"`user_id`",
		"`vote_count`",
		"`answer_count`",
		"0 as `accepted`",
		"`status`",
		"`post_update_time`",
	).From("question")

	for i, word := range words {
		if i == 0 {
			b.Where(builder.Like{"original_text", word})
		} else {
			b.Or(builder.Like{"original_text", word})
		}
	}

	// check need filter has not accepted
	if limitNoAccepted {
		b.And(builder.Eq{"accepted_answer_id": 0})
	}

	if answers == 0 {
		b.And(builder.Eq{"answer_count": 0})
	} else if answers > 0 {
		b.And(builder.Gte{"answer_count": answers})
	}
	res, err := sr.data.DB.Query(b.OrderBy(sr.parseOrder(ctx, order)).Limit(size, page-1))

	tr, err := sr.data.DB.Query(builder.MySQL().Select("count(*) total").From(b, "c"))

	if len(tr) != 0 {
		total = converter.StringToInt64(string(tr[0]["total"]))
	}
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
		return
	}
	resp, err = sr.parseResult(ctx, res)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

func (sr *searchRepo) SearchAnswers(ctx context.Context, words []string, limitAccepted bool, questionID string, page, size int, order string) (resp []schema.SearchResp, total int64, err error) {
	b := builder.MySQL().Select(
		"`answer`.`id` as `id`",
		"`question_id`",
		"`question`.`title` as `title`",
		"`answer`.`original_text` as `original_text`",
		"`answer`.`created_at`",
		"`answer`.`user_id` as `user_id`",
		"`answer`.`vote_count` as `vote_count`",
		"0 as `answer_count`",
		"`adopted` as `accepted`",
		"`answer`.`status` as `status`",
		"`answer`.`created_at` as `post_update_time`",
	).From("`answer`").
		LeftJoin("`question`", "`question`.id = `answer`.question_id")

	for i, word := range words {
		if i == 0 {
			b.Where(builder.Like{"`answer`.original_text", word})
		} else {
			b.Or(builder.Like{"`answer`.original_text", word})
		}
	}

	if limitAccepted {
		b.Where(builder.Eq{"adopted": 2})
	}

	if questionID != "" {
		b.Where(builder.Eq{"question_id": questionID})
	}

	res, err := sr.data.DB.Query(b.OrderBy(sr.parseOrder(ctx, order)).Limit(size, page-1))

	tr, err := sr.data.DB.Query(builder.MySQL().Select("count(*) total").From(b, "c"))
	total = converter.StringToInt64(string(tr[0]["total"]))
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
		return
	}
	resp, err = sr.parseResult(ctx, res)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

func (sr *searchRepo) parseOrder(ctx context.Context, order string) (res string) {
	switch order {
	case "newest":
		res = "created_at desc"
	case "active":
		res = "post_update_time desc"
	case "score":
		res = "vote_count desc"
	default:
		res = "created_at desc"
	}
	return
}

func (sr *searchRepo) parseResult(ctx context.Context, res []map[string][]byte) (resp []schema.SearchResp, err error) {
	for _, r := range res {
		var (
			objectKey,
			status string

			tags       []schema.TagResp
			tagsEntity []entity.Tag
			object     schema.SearchObject
		)
		objectKey, err = obj.GetObjectTypeStrByObjectID(string(r["id"]))
		if err != nil {
			continue
		}

		tp, _ := time.ParseInLocation("2006-01-02 15:04:05", string(r["created_at"]), time.Local)

		// get user info
		userInfo, _, e := sr.userCommon.GetUserBasicInfoByID(ctx, string(r["user_id"]))
		if e != nil {
			err = errors.InternalServer(reason.DatabaseError).WithError(e).WithStack()
			return
		}

		// get tags
		err = sr.data.DB.
			Select("`display_name`,`slug_name`,`main_tag_slug_name`").
			Table("tag").
			Join("INNER", "tag_rel", "tag.id = tag_rel.tag_id").
			Where(builder.Eq{"tag_rel.object_id": r["question_id"]}).
			And(builder.Eq{"tag_rel.status": entity.TagRelStatusAvailable}).
			Find(&tagsEntity)

		if err != nil {
			err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
			return
		}
		_ = copier.Copy(&tags, tagsEntity)
		switch objectKey {
		case "question":
			for k, v := range entity.CmsQuestionSearchStatus {
				if v == converter.StringToInt(string(r["status"])) {
					status = k
					break
				}
			}
		case "answer":
			for k, v := range entity.CmsAnswerSearchStatus {
				if v == converter.StringToInt(string(r["status"])) {
					status = k
					break
				}
			}
		}

		object = schema.SearchObject{
			ID:              string(r["id"]),
			Title:           string(r["title"]),
			Excerpt:         cutOutParsedText(string(r["original_text"])),
			CreatedAtParsed: tp.Unix(),
			UserInfo:        userInfo,
			Tags:            tags,
			VoteCount:       converter.StringToInt(string(r["vote_count"])),
			Accepted:        string(r["accepted"]) == "2",
			AnswerCount:     converter.StringToInt(string(r["answer_count"])),
			StatusStr:       status,
		}
		resp = append(resp, schema.SearchResp{
			ObjectType: objectKey,
			Object:     object,
		})
	}
	return
}

// userBasicInfoFormat
func (sr *searchRepo) userBasicInfoFormat(ctx context.Context, dbinfo *entity.User) *schema.UserBasicInfo {
	return &schema.UserBasicInfo{
		ID:          dbinfo.ID,
		Username:    dbinfo.Username,
		Rank:        dbinfo.Rank,
		DisplayName: dbinfo.DisplayName,
		Avatar:      dbinfo.Avatar,
		Website:     dbinfo.Website,
		Location:    dbinfo.Location,
		IpInfo:      dbinfo.IPInfo,
	}
}

func cutOutParsedText(parsedText string) string {
	parsedText = strings.TrimSpace(parsedText)
	idx := strings.Index(parsedText, "\n")
	if idx >= 0 {
		parsedText = parsedText[0:idx]
	}
	return parsedText
}

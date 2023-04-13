package migrations

import (
	"time"

	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/service/permission"
	"xorm.io/xorm"
)

func addRolePinAndHideFeatures(x *xorm.Engine) error {

	powers := []*entity.Power{
		{ID: 34, Name: "question pin", PowerType: permission.QuestionPin, Description: "Top or untop the question"},
		{ID: 35, Name: "question hide", PowerType: permission.QuestionHide, Description: "hide or show the question"},
	}
	// insert default powers
	for _, power := range powers {
		exist, err := x.Get(&entity.Power{ID: power.ID})
		if err != nil {
			return err
		}
		if exist {
			_, err = x.ID(power.ID).Update(power)
		} else {
			_, err = x.Insert(power)
		}
		if err != nil {
			return err
		}
	}

	rolePowerRels := []*entity.RolePowerRel{

		{RoleID: 2, PowerType: permission.QuestionPin},
		{RoleID: 2, PowerType: permission.QuestionHide},

		{RoleID: 3, PowerType: permission.QuestionPin},
		{RoleID: 3, PowerType: permission.QuestionHide},
	}

	// insert default powers
	for _, rel := range rolePowerRels {
		exist, err := x.Get(&entity.RolePowerRel{RoleID: rel.RoleID, PowerType: rel.PowerType})
		if err != nil {
			return err
		}
		if exist {
			continue
		}
		_, err = x.Insert(rel)
		if err != nil {
			return err
		}
	}
	type Question struct {
		ID               string    `xorm:"not null pk BIGINT(20) id"`
		CreatedAt        time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP created_at"`
		UpdatedAt        time.Time `xorm:"updated_at TIMESTAMP"`
		UserID           string    `xorm:"not null default 0 BIGINT(20) INDEX user_id"`
		LastEditUserID   string    `xorm:"not null default 0 BIGINT(20) last_edit_user_id"`
		Title            string    `xorm:"not null default '' VARCHAR(150) title"`
		OriginalText     string    `xorm:"not null MEDIUMTEXT original_text"`
		ParsedText       string    `xorm:"not null MEDIUMTEXT parsed_text"`
		Status           int       `xorm:"not null default 1 INT(11) status"`
		Pin              int       `xorm:"not null default 1 INT(11) pin"`
		Show             int       `xorm:"not null default 1 INT(11) show"`
		ViewCount        int       `xorm:"not null default 0 INT(11) view_count"`
		UniqueViewCount  int       `xorm:"not null default 0 INT(11) unique_view_count"`
		VoteCount        int       `xorm:"not null default 0 INT(11) vote_count"`
		AnswerCount      int       `xorm:"not null default 0 INT(11) answer_count"`
		CollectionCount  int       `xorm:"not null default 0 INT(11) collection_count"`
		FollowCount      int       `xorm:"not null default 0 INT(11) follow_count"`
		AcceptedAnswerID string    `xorm:"not null default 0 BIGINT(20) accepted_answer_id"`
		LastAnswerID     string    `xorm:"not null default 0 BIGINT(20) last_answer_id"`
		PostUpdateTime   time.Time `xorm:"post_update_time TIMESTAMP"`
		RevisionID       string    `xorm:"not null default 0 BIGINT(20) revision_id"`
	}
	err := x.Sync(new(Question))
	if err != nil {
		return err
	}

	return nil
}

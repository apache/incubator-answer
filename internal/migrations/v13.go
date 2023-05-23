package migrations

import (
	"encoding/json"
	"fmt"

	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/permission"
	"xorm.io/xorm"
)

func updateCount(x *xorm.Engine) error {
	addPrivilegeForInviteSomeoneToAnswer(x)
	addGravatarBaseURL(x)
	return nil
}

func addGravatarBaseURL(x *xorm.Engine) error {
	usersSiteInfo := &entity.SiteInfo{
		Type: constant.SiteTypeUsers,
	}
	exist, err := x.Get(usersSiteInfo)
	if err != nil {
		return fmt.Errorf("get config failed: %w", err)
	}
	if exist {
		content := &schema.SiteUsersReq{}
		_ = json.Unmarshal([]byte(usersSiteInfo.Content), content)
		content.GravatarBaseURL = "https://www.gravatar.com/avatar/"
		data, _ := json.Marshal(content)
		usersSiteInfo.Content = string(data)

		_, err = x.ID(usersSiteInfo.ID).Cols("content").Update(usersSiteInfo)
		if err != nil {
			return fmt.Errorf("update site info failed: %w", err)
		}
	}
	return nil
}

func addPrivilegeForInviteSomeoneToAnswer(x *xorm.Engine) error {
	// add rank for invite to answer
	powers := []*entity.Power{
		{ID: 38, Name: "invite someone to answer", PowerType: permission.AnswerInviteSomeoneToAnswer, Description: "invite someone to answer"},
	}
	for _, power := range powers {
		exist, err := x.Get(&entity.Power{PowerType: power.PowerType})
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
		{RoleID: 2, PowerType: permission.AnswerInviteSomeoneToAnswer},
		{RoleID: 3, PowerType: permission.AnswerInviteSomeoneToAnswer},
	}
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

	defaultConfigTable := []*entity.Config{
		{ID: 127, Key: "rank.answer.invite_someone_to_answer", Value: `1000`},
	}
	for _, c := range defaultConfigTable {
		exist, err := x.Get(&entity.Config{ID: c.ID})
		if err != nil {
			return fmt.Errorf("get config failed: %w", err)
		}
		if exist {
			if _, err = x.Update(c, &entity.Config{ID: c.ID}); err != nil {
				return fmt.Errorf("update config failed: %w", err)
			}
			continue
		}
		if _, err = x.Insert(&entity.Config{ID: c.ID, Key: c.Key, Value: c.Value}); err != nil {
			return fmt.Errorf("add config failed: %w", err)
		}
	}
	return nil
}

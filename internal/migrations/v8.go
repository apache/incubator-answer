package migrations

import (
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

	return nil
}

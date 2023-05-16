package migrations

import (
	"encoding/json"
	"fmt"

	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
	"xorm.io/xorm"
)

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

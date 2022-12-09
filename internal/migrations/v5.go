package migrations

import (
	"encoding/json"

	"github.com/answerdev/answer/internal/entity"
	"xorm.io/xorm"
)

func addThemeAndPrivateMode(x *xorm.Engine) error {
	loginConfig := map[string]bool{
		"allow_new_registrations": true,
		"login_required":          false,
	}
	loginConfigDataBytes, _ := json.Marshal(loginConfig)
	_, err := x.InsertOne(&entity.SiteInfo{
		Type:    "login",
		Content: string(loginConfigDataBytes),
		Status:  1,
	})
	return err
}

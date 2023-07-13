package migrations

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/answerdev/answer/internal/entity"
	"xorm.io/xorm"
)

func addThemeAndPrivateMode(ctx context.Context, x *xorm.Engine) error {
	loginConfig := map[string]bool{
		"allow_new_registrations": true,
		"login_required":          false,
	}
	loginConfigDataBytes, _ := json.Marshal(loginConfig)
	siteInfo := &entity.SiteInfo{
		Type:    "login",
		Content: string(loginConfigDataBytes),
		Status:  1,
	}
	exist, err := x.Context(ctx).Get(&entity.SiteInfo{Type: siteInfo.Type})
	if err != nil {
		return fmt.Errorf("get config failed: %w", err)
	}
	if !exist {
		_, err = x.Context(ctx).Insert(siteInfo)
		if err != nil {
			return fmt.Errorf("insert site info failed: %w", err)
		}
	}

	themeConfig := `{"theme":"default","theme_config":{"default":{"navbar_style":"colored","primary_color":"#0033ff"}}}`
	themeSiteInfo := &entity.SiteInfo{
		Type:    "theme",
		Content: themeConfig,
		Status:  1,
	}
	exist, err = x.Context(ctx).Get(&entity.SiteInfo{Type: themeSiteInfo.Type})
	if err != nil {
		return fmt.Errorf("get config failed: %w", err)
	}
	if !exist {
		_, err = x.Context(ctx).Insert(themeSiteInfo)
	}
	return err
}

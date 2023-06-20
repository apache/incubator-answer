package migrations

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/answerdev/answer/internal/entity"
	"golang.org/x/crypto/bcrypt"
	"xorm.io/xorm"
)

type Mentor struct {
	ctx      context.Context
	engine   *xorm.Engine
	userData *InitNeedUserInputData
	err      error
	Done     bool
}

func NewMentor(ctx context.Context, engine *xorm.Engine, data *InitNeedUserInputData) *Mentor {
	return &Mentor{ctx: ctx, engine: engine, userData: data}
}

type InitNeedUserInputData struct {
	Language      string
	SiteName      string
	SiteURL       string
	ContactEmail  string
	AdminName     string
	AdminPassword string
	AdminEmail    string
}

func (m *Mentor) InitDB() error {
	m.do("check table exist", m.checkTableExist)
	m.do("sync table", m.syncTable)
	m.do("init version table", m.initVersionTable)
	m.do("init admin user", m.initAdminUser)
	m.do("init config", m.initConfig)
	m.do("init role", m.initRole)
	m.do("init power", m.initPower)
	m.do("init role power rel", m.initRolePowerRel)
	m.do("init admin user role rel", m.initAdminUserRoleRel)
	m.do("init site info interface", m.initSiteInfoInterface)
	m.do("init site info general config", m.initSiteInfoGeneralData)
	m.do("init site info login config", m.initSiteInfoLoginConfig)
	m.do("init site info theme config", m.initSiteInfoThemeConfig)
	m.do("init site info seo config", m.initSiteInfoSEOConfig)
	m.do("init site info user config", m.initSiteInfoUsersConfig)
	return m.err
}

func (m *Mentor) do(taskName string, fn func()) {
	if m.err != nil || m.Done {
		return
	}
	fn()
	if m.err != nil {
		m.err = fmt.Errorf("%s failed: %s", taskName, m.err)
	}
}

func (m *Mentor) checkTableExist() {
	m.Done, m.err = m.engine.Context(m.ctx).IsTableExist(&entity.Version{})
	if m.Done {
		fmt.Println("[database] already exists")
	}
}

func (m *Mentor) syncTable() {
	m.err = m.engine.Context(m.ctx).Sync(tables...)
}

func (m *Mentor) initVersionTable() {
	_, m.err = m.engine.Context(m.ctx).Insert(&entity.Version{ID: 1, VersionNumber: ExpectedVersion()})
}

func (m *Mentor) initAdminUser() {
	generateFromPassword, _ := bcrypt.GenerateFromPassword([]byte(m.userData.AdminPassword), bcrypt.DefaultCost)
	_, m.err = m.engine.Context(m.ctx).Insert(&entity.User{
		ID:           "1",
		Username:     m.userData.AdminName,
		Pass:         string(generateFromPassword),
		EMail:        m.userData.AdminEmail,
		MailStatus:   1,
		NoticeStatus: 1,
		Status:       1,
		Rank:         1,
		DisplayName:  m.userData.AdminName,
	})
}

func (m *Mentor) initConfig() {
	_, m.err = m.engine.Context(m.ctx).Insert(defaultConfigTable)
}

func (m *Mentor) initRole() {
	_, m.err = m.engine.Context(m.ctx).Insert(roles)
}

func (m *Mentor) initPower() {
	_, m.err = m.engine.Context(m.ctx).Insert(powers)
}

func (m *Mentor) initRolePowerRel() {
	_, m.err = m.engine.Context(m.ctx).Insert(rolePowerRels)
}

func (m *Mentor) initAdminUserRoleRel() {
	_, m.err = m.engine.Context(m.ctx).Insert(adminUserRoleRel)
}

func (m *Mentor) initSiteInfoInterface() {
	interfaceData := map[string]string{
		"language":  m.userData.Language,
		"time_zone": "UTC",
	}
	interfaceDataBytes, _ := json.Marshal(interfaceData)
	_, m.err = m.engine.Context(m.ctx).Insert(&entity.SiteInfo{
		Type:    "interface",
		Content: string(interfaceDataBytes),
		Status:  1,
	})
}

func (m *Mentor) initSiteInfoGeneralData() {
	generalData := map[string]string{
		"name":          m.userData.SiteName,
		"site_url":      m.userData.SiteURL,
		"contact_email": m.userData.ContactEmail,
	}
	generalDataBytes, _ := json.Marshal(generalData)
	_, m.err = m.engine.Context(m.ctx).Insert(&entity.SiteInfo{
		Type:    "general",
		Content: string(generalDataBytes),
		Status:  1,
	})
}

func (m *Mentor) initSiteInfoLoginConfig() {
	loginConfig := map[string]bool{
		"allow_new_registrations":   true,
		"allow_email_registrations": true,
		"login_required":            false,
	}
	loginConfigDataBytes, _ := json.Marshal(loginConfig)
	_, m.err = m.engine.Context(m.ctx).Insert(&entity.SiteInfo{
		Type:    "login",
		Content: string(loginConfigDataBytes),
		Status:  1,
	})
}

func (m *Mentor) initSiteInfoThemeConfig() {
	themeConfig := `{"theme":"default","theme_config":{"default":{"navbar_style":"colored","primary_color":"#0033ff"}}}`
	_, m.err = m.engine.Context(m.ctx).Insert(&entity.SiteInfo{
		Type:    "theme",
		Content: themeConfig,
		Status:  1,
	})
}

func (m *Mentor) initSiteInfoSEOConfig() {
	seoData := map[string]interface{}{
		"permalink": 1,
		"robots":    defaultSEORobotTxt + m.userData.SiteURL + "/sitemap.xml",
	}
	seoDataBytes, _ := json.Marshal(seoData)
	_, m.err = m.engine.Context(m.ctx).Insert(&entity.SiteInfo{
		Type:    "seo",
		Content: string(seoDataBytes),
		Status:  1,
	})
}

func (m *Mentor) initSiteInfoUsersConfig() {
	usersData := map[string]any{
		"default_avatar":            "gravatar",
		"gravatar_base_url":         "https://www.gravatar.com/avatar/",
		"allow_update_display_name": true,
		"allow_update_username":     true,
		"allow_update_avatar":       true,
		"allow_update_bio":          true,
		"allow_update_website":      true,
		"allow_update_location":     true,
	}
	usersDataBytes, _ := json.Marshal(usersData)
	_, m.err = m.engine.Context(m.ctx).Insert(&entity.SiteInfo{
		Type:    "users",
		Content: string(usersDataBytes),
		Status:  1,
	})
}

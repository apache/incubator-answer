package migrations

import (
	"encoding/json"
	"fmt"

	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/entity"
	"golang.org/x/crypto/bcrypt"
	"xorm.io/xorm"
)

var tables = []interface{}{
	&entity.Activity{},
	&entity.Answer{},
	&entity.Collection{},
	&entity.CollectionGroup{},
	&entity.Comment{},
	&entity.Config{},
	&entity.Meta{},
	&entity.Notification{},
	&entity.Question{},
	&entity.Report{},
	&entity.Revision{},
	&entity.SiteInfo{},
	&entity.Tag{},
	&entity.TagRel{},
	&entity.Uniqid{},
	&entity.User{},
	&entity.Version{},
}

// InitDB init db
func InitDB(dataConf *data.Database) (err error) {
	engine, err := data.NewDB(false, dataConf)
	if err != nil {
		fmt.Println("new database failed: ", err.Error())
		return err
	}

	exist, err := engine.IsTableExist(&entity.Version{})
	if err != nil {
		return fmt.Errorf("check table exists failed: %s", err)
	}
	if exist {
		fmt.Println("[database] already exists")
		return nil
	}

	err = engine.Sync(tables...)
	if err != nil {
		return fmt.Errorf("sync table failed: %s", err)
	}

	err = initAdminUser(engine)
	if err != nil {
		return fmt.Errorf("init admin user failed: %s", err)
	}

	err = initConfigTable(engine)
	if err != nil {
		return fmt.Errorf("init config table: %s", err)
	}
	return nil
}

func initAdminUser(engine *xorm.Engine) error {
	_, err := engine.InsertOne(&entity.User{
		Username:     "admin",
		Pass:         "$2a$10$.gnUnpW.8ssRNaEvx.XwvOR2NuPsGzFLWWX2rqSIVAdIvLNZZYs5y", // admin
		EMail:        "admin@admin.com",
		MailStatus:   1,
		NoticeStatus: 1,
		Status:       1,
		Rank:         1,
		DisplayName:  "admin",
		IsAdmin:      true,
	})
	return err
}

func initSiteInfo(engine *xorm.Engine, language, siteName, siteURL, contactEmail string) error {
	interfaceData := map[string]string{
		"logo":     "",
		"theme":    "black",
		"language": language,
	}
	interfaceDataBytes, _ := json.Marshal(interfaceData)
	_, err := engine.InsertOne(&entity.SiteInfo{
		Type:    "interface",
		Content: string(interfaceDataBytes),
		Status:  1,
	})
	if err != nil {
		return err
	}

	generalData := map[string]string{
		"name":          siteName,
		"site_url":      siteURL,
		"contact_email": contactEmail,
	}
	generalDataBytes, _ := json.Marshal(generalData)
	_, err = engine.InsertOne(&entity.SiteInfo{
		Type:    "general",
		Content: string(generalDataBytes),
		Status:  1,
	})
	return err
}

func updateAdminInfo(engine *xorm.Engine, adminName, adminPassword, adminEmail string) error {
	generateFromPassword, err := bcrypt.GenerateFromPassword([]byte(adminPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("")
	}
	adminPassword = string(generateFromPassword)

	// update admin info
	_, err = engine.ID("1").Update(&entity.User{
		Username:    adminName,
		Pass:        adminPassword,
		EMail:       adminEmail,
		DisplayName: adminName,
	})
	if err != nil {
		return fmt.Errorf("update admin user info failed: %s", err)
	}
	return nil
}

// UpdateInstallInfo update some init data about the admin interface and admin password
func UpdateInstallInfo(dataConf *data.Database, language string,
	siteName string,
	siteURL string,
	contactEmail string,
	adminName string,
	adminPassword string,
	adminEmail string) error {

	engine, err := data.NewDB(false, dataConf)
	if err != nil {
		return fmt.Errorf("database connection error: %s", err)
	}

	err = updateAdminInfo(engine, adminName, adminPassword, adminEmail)
	if err != nil {
		return fmt.Errorf("update admin info failed: %s", err)
	}

	err = initSiteInfo(engine, language, siteName, siteURL, contactEmail)
	if err != nil {
		return fmt.Errorf("init site info failed: %s", err)
	}
	return err
}

func initConfigTable(engine *xorm.Engine) error {
	defaultConfigTable := []*entity.Config{
		{ID: 1, Key: "answer.accepted", Value: `15`},
		{ID: 2, Key: "answer.voted_up", Value: `10`},
		{ID: 3, Key: "question.voted_up", Value: `10`},
		{ID: 4, Key: "tag.edit_accepted", Value: `2`},
		{ID: 5, Key: "answer.accept", Value: `2`},
		{ID: 6, Key: "answer.voted_down_cancel", Value: `2`},
		{ID: 7, Key: "question.voted_down_cancel", Value: `2`},
		{ID: 8, Key: "answer.vote_down_cancel", Value: `1`},
		{ID: 9, Key: "question.vote_down_cancel", Value: `1`},
		{ID: 10, Key: "user.activated", Value: `1`},
		{ID: 11, Key: "edit.accepted", Value: `2`},
		{ID: 12, Key: "answer.vote_down", Value: `-1`},
		{ID: 13, Key: "question.voted_down", Value: `-2`},
		{ID: 14, Key: "answer.voted_down", Value: `-2`},
		{ID: 15, Key: "answer.accept_cancel", Value: `-2`},
		{ID: 16, Key: "answer.deleted", Value: `-5`},
		{ID: 17, Key: "question.voted_up_cancel", Value: `-10`},
		{ID: 18, Key: "answer.voted_up_cancel", Value: `-10`},
		{ID: 19, Key: "answer.accepted_cancel", Value: `-15`},
		{ID: 20, Key: "object.reported", Value: `-100`},
		{ID: 21, Key: "edit.rejected", Value: `-2`},
		{ID: 22, Key: "daily_rank_limit", Value: `200`},
		{ID: 23, Key: "daily_rank_limit.exclude", Value: `["answer.accepted"]`},
		{ID: 24, Key: "user.follow", Value: `0`},
		{ID: 25, Key: "comment.vote_up", Value: `0`},
		{ID: 26, Key: "comment.vote_up_cancel", Value: `0`},
		{ID: 27, Key: "question.vote_down", Value: `0`},
		{ID: 28, Key: "question.vote_up", Value: `0`},
		{ID: 29, Key: "question.vote_up_cancel", Value: `0`},
		{ID: 30, Key: "answer.vote_up", Value: `0`},
		{ID: 31, Key: "answer.vote_up_cancel", Value: `0`},
		{ID: 32, Key: "question.follow", Value: `0`},
		{ID: 33, Key: "email.config", Value: `{"from_name":"answer","from_email":"answer@answer.com","smtp_host":"smtp.answer.org","smtp_port":465,"smtp_password":"answer","smtp_username":"answer@answer.com","smtp_authentication":true,"encryption":"","register_title":"[{{.SiteName}}] Confirm your new account","register_body":"Welcome to {{.SiteName}}<br><br>\n\nClick the following link to confirm and activate your new account:<br>\n<a href='{{.RegisterUrl}}' target='_blank'>{{.RegisterUrl}}</a><br><br>\n\nIf the above link is not clickable, try copying and pasting it into the address bar of your web browser.\n","pass_reset_title":"[{{.SiteName }}] Password reset","pass_reset_body":"Somebody asked to reset your password on [{{.SiteName}}].<br><br>\n\nIf it was not you, you can safely ignore this email.<br><br>\n\nClick the following link to choose a new password:<br>\n<a href='{{.PassResetUrl}}' target='_blank'>{{.PassResetUrl}}</a>\n","change_title":"[{{.SiteName}}] Confirm your new email address","change_body":"Confirm your new email address for {{.SiteName}}  by clicking on the following link:<br><br>\n\n<a href='{{.ChangeEmailUrl}}' target='_blank'>{{.ChangeEmailUrl}}</a><br><br>\n\nIf you did not request this change, please ignore this email.\n","test_title":"[{{.SiteName}}] Test Email","test_body":"This is a test email."}`},
		{ID: 35, Key: "tag.follow", Value: `0`},
		{ID: 36, Key: "rank.question.add", Value: `0`},
		{ID: 37, Key: "rank.question.edit", Value: `0`},
		{ID: 38, Key: "rank.question.delete", Value: `0`},
		{ID: 39, Key: "rank.question.vote_up", Value: `0`},
		{ID: 40, Key: "rank.question.vote_down", Value: `0`},
		{ID: 41, Key: "rank.answer.add", Value: `0`},
		{ID: 42, Key: "rank.answer.edit", Value: `0`},
		{ID: 43, Key: "rank.answer.delete", Value: `0`},
		{ID: 44, Key: "rank.answer.accept", Value: `0`},
		{ID: 45, Key: "rank.answer.vote_up", Value: `0`},
		{ID: 46, Key: "rank.answer.vote_down", Value: `0`},
		{ID: 47, Key: "rank.comment.add", Value: `0`},
		{ID: 48, Key: "rank.comment.edit", Value: `0`},
		{ID: 49, Key: "rank.comment.delete", Value: `0`},
		{ID: 50, Key: "rank.report.add", Value: `0`},
		{ID: 51, Key: "rank.tag.add", Value: `0`},
		{ID: 52, Key: "rank.tag.edit", Value: `0`},
		{ID: 53, Key: "rank.tag.delete", Value: `0`},
		{ID: 54, Key: "rank.tag.synonym", Value: `0`},
		{ID: 55, Key: "rank.link.url_limit", Value: `0`},
		{ID: 56, Key: "rank.vote.detail", Value: `0`},
		{ID: 57, Key: "reason.spam", Value: `{"name":"spam","description":"This post is an advertisement, or vandalism. It is not useful or relevant to the current topic."}`},
		{ID: 58, Key: "reason.rude_or_abusive", Value: `{"name":"rude or abusive","description":"A reasonable person would find this content inappropriate for respectful discourse."}`},
		{ID: 59, Key: "reason.something", Value: `{"name":"something else","description":"This post requires staff attention for another reason not listed above.","content_type":"textarea"}`},
		{ID: 60, Key: "reason.a_duplicate", Value: `{"name":"a duplicate","description":"This question has been asked before and already has an answer.","content_type":"text"}`},
		{ID: 61, Key: "reason.not_a_answer", Value: `{"name":"not a answer","description":"This was posted as an answer, but it does not attempt to answer the question. It should possibly be an edit, a comment, another question, or deleted altogether.","content_type":""}`},
		{ID: 62, Key: "reason.no_longer_needed", Value: `{"name":"no longer needed","description":"This comment is outdated, conversational or not relevant to this post."}`},
		{ID: 63, Key: "reason.community_specific", Value: `{"name":"a community-specific reason","description":"This question doesn’t meet a community guideline."}`},
		{ID: 64, Key: "reason.not_clarity", Value: `{"name":"needs details or clarity","description":"This question currently includes multiple questions in one. It should focus on one problem only.","content_type":"text"}`},
		{ID: 65, Key: "reason.normal", Value: `{"name":"normal","description":"A normal post available to everyone."}`},
		{ID: 66, Key: "reason.normal.user", Value: `{"name":"normal","description":"A normal user can ask and answer questions."}`},
		{ID: 67, Key: "reason.closed", Value: `{"name":"closed","description":"A closed question can’t answer, but still can edit, vote and comment."}`},
		{ID: 68, Key: "reason.deleted", Value: `{"name":"deleted","description":"All reputation gained and lost will be restored."}`},
		{ID: 69, Key: "reason.deleted.user", Value: `{"name":"deleted","description":"Delete profile, authentication associations."}`},
		{ID: 70, Key: "reason.suspended", Value: `{"name":"suspended","description":"A suspended user can’t log in."}`},
		{ID: 71, Key: "reason.inactive", Value: `{"name":"inactive","description":"An inactive user must re-validate their email."}`},
		{ID: 72, Key: "reason.looks_ok", Value: `{"name":"looks ok","description":"This post is good as-is and not low quality."}`},
		{ID: 73, Key: "reason.needs_edit", Value: `{"name":"needs edit, and I did it","description":"Improve and correct problems with this post yourself."}`},
		{ID: 74, Key: "reason.needs_close", Value: `{"name":"needs close","description":"A closed question can’t answer, but still can edit, vote and comment."}`},
		{ID: 75, Key: "reason.needs_delete", Value: `{"name":"needs delete","description":"All reputation gained and lost will be restored."}`},
		{ID: 76, Key: "question.flag.reasons", Value: `["reason.spam","reason.rude_or_abusive","reason.something","reason.a_duplicate"]`},
		{ID: 77, Key: "answer.flag.reasons", Value: `["reason.spam","reason.rude_or_abusive","reason.something","reason.not_a_answer"]`},
		{ID: 78, Key: "comment.flag.reasons", Value: `["reason.spam","reason.rude_or_abusive","reason.something","reason.no_longer_needed"]`},
		{ID: 79, Key: "question.close.reasons", Value: `["reason.a_duplicate","reason.community_specific","reason.not_clarity","reason.something"]`},
		{ID: 80, Key: "question.status.reasons", Value: `["reason.normal","reason.closed","reason.deleted"]`},
		{ID: 81, Key: "answer.status.reasons", Value: `["reason.normal","reason.deleted"]`},
		{ID: 82, Key: "comment.status.reasons", Value: `["reason.normal","reason.deleted"]`},
		{ID: 83, Key: "user.status.reasons", Value: `["reason.normal.user","reason.suspended","reason.deleted.user","reason.inactive"]`},
		{ID: 84, Key: "question.review.reasons", Value: `["reason.looks_ok","reason.needs_edit","reason.needs_close","reason.needs_delete"]`},
		{ID: 85, Key: "answer.review.reasons", Value: `["reason.looks_ok","reason.needs_edit","reason.needs_delete"]`},
		{ID: 86, Key: "comment.review.reasons", Value: `["reason.looks_ok","reason.needs_edit","reason.needs_delete"]`},
	}
	_, err := engine.Insert(defaultConfigTable)
	return err
}

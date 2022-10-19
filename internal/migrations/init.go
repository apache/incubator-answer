package migrations

import (
	"fmt"

	"github.com/segmentfault/answer/internal/base/data"
	"github.com/segmentfault/answer/internal/entity"
	"xorm.io/xorm"
)

var (
	tables = []interface{}{
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
)

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

	err = initSiteInfo(engine)
	if err != nil {
		return fmt.Errorf("init site info failed: %s", err)
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
		DisplayName:  "admin",
		IsAdmin:      true,
	})
	return err
}

func initSiteInfo(engine *xorm.Engine) error {
	_, err := engine.InsertOne(&entity.SiteInfo{
		Type:    "interface",
		Content: `{"logo":"","theme":"black","language":"en_US"}`,
		Status:  1,
	})
	return err
}

func initConfigTable(engine *xorm.Engine) error {
	defaultConfigTable := []*entity.Config{
		{1, "answer.accepted", `15`},
		{2, "answer.voted_up", `10`},
		{3, "question.voted_up", `10`},
		{4, "tag.edit_accepted", `2`},
		{5, "answer.accept", `2`},
		{6, "answer.voted_down_cancel", `2`},
		{7, "question.voted_down_cancel", `2`},
		{8, "answer.vote_down_cancel", `1`},
		{9, "question.vote_down_cancel", `1`},
		{10, "user.activated", `1`},
		{11, "edit.accepted", `2`},
		{12, "answer.vote_down", `-1`},
		{13, "question.voted_down", `-2`},
		{14, "answer.voted_down", `-2`},
		{15, "answer.accept_cancel", `-2`},
		{16, "answer.deleted", `-5`},
		{17, "question.voted_up_cancel", `-10`},
		{18, "answer.voted_up_cancel", `-10`},
		{19, "answer.accepted_cancel", `-15`},
		{20, "object.reported", `-100`},
		{21, "edit.rejected", `-2`},
		{22, "daily_rank_limit", `200`},
		{23, "daily_rank_limit.exclude", `["answer.accepted"]`},
		{24, "user.follow", `0`},
		{25, "comment.vote_up", `0`},
		{26, "comment.vote_up_cancel", `0`},
		{27, "question.vote_down", `0`},
		{28, "question.vote_up", `0`},
		{29, "question.vote_up_cancel", `0`},
		{30, "answer.vote_up", `0`},
		{31, "answer.vote_up_cancel", `0`},
		{32, "question.follow", `0`},
		{33, "email.config", `{"email_web_name":"answer","email_from":"","email_from_pass":"","email_from_hostname":"","email_from_smtp":"","email_from_name":"Answer Team","email_register_title":"[{{.SiteName}}] Confirm your new account","email_register_body":"Welcome to {{.SiteName}}<br><br>\\n\\nClick the following link to confirm and activate your new account:<br>\\n{{.RegisterUrl}}<br><br>\\n\\nIf the above link is not clickable, try copying and pasting it into the address bar of your web browser.\\n","email_pass_reset_title":"[{{.SiteName }}] Password reset","email_pass_reset_body":"Somebody asked to reset your password on [{{.SiteName}}].<br><br>\\n\\nIf it was not you, you can safely ignore this email.<br><br>\\n\\nClick the following link to choose a new password:<br>\\n{{.PassResetUrl}}\\n","email_change_title":"[{{.SiteName}}] Confirm your new email address","email_change_body":"Confirm your new email address for {{.SiteName}}  by clicking on the following link:<br><br>\\n\\n{{.ChangeEmailUrl}}<br><br>\\n\\nIf you did not request this change, please ignore this email.\\n"}`},
		{35, "tag.follow", `0`},
		{36, "rank.question.add", `0`},
		{37, "rank.question.edit", `0`},
		{38, "rank.question.delete", `0`},
		{39, "rank.question.vote_up", `0`},
		{40, "rank.question.vote_down", `0`},
		{41, "rank.answer.add", `0`},
		{42, "rank.answer.edit", `0`},
		{43, "rank.answer.delete", `0`},
		{44, "rank.answer.accept", `0`},
		{45, "rank.answer.vote_up", `0`},
		{46, "rank.answer.vote_down", `0`},
		{47, "rank.comment.add", `0`},
		{48, "rank.comment.edit", `0`},
		{49, "rank.comment.delete", `0`},
		{50, "rank.report.add", `0`},
		{51, "rank.tag.add", `0`},
		{52, "rank.tag.edit", `0`},
		{53, "rank.tag.delete", `0`},
		{54, "rank.tag.synonym", `0`},
		{55, "rank.link.url_limit", `0`},
		{56, "rank.vote.detail", `0`},
		{57, "reason.spam", `{"name":"spam","description":"This post is an advertisement, or vandalism. It is not useful or relevant to the current topic."}`},
		{58, "reason.rude_or_abusive", `{"name":"rude or abusive","description":"A reasonable person would find this content inappropriate for respectful discourse."}`},
		{59, "reason.something", `{"name":"something else","description":"This post requires staff attention for another reason not listed above.","content_type":"textarea"}`},
		{60, "reason.a_duplicate", `{"name":"a duplicate","description":"This question has been asked before and already has an answer.","content_type":"text"}`},
		{61, "reason.not_a_answer", `{"name":"not a answer","description":"This was posted as an answer, but it does not attempt to answer the question. It should possibly be an edit, a comment, another question, or deleted altogether.","content_type":""}`},
		{62, "reason.no_longer_needed", `{"name":"no longer needed","description":"This comment is outdated, conversational or not relevant to this post."}`},
		{63, "reason.community_specific", `{"name":"a community-specific reason","description":"This question doesn’t meet a community guideline."}`},
		{64, "reason.not_clarity", `{"name":"needs details or clarity","description":"This question currently includes multiple questions in one. It should focus on one problem only.","content_type":"text"}`},
		{65, "reason.normal", `{"name":"normal","description":"A normal post available to everyone."}`},
		{66, "reason.normal.user", `{"name":"normal","description":"A normal user can ask and answer questions."}`},
		{67, "reason.closed", `{"name":"closed","description":"A closed question can’t answer, but still can edit, vote and comment."}`},
		{68, "reason.deleted", `{"name":"deleted","description":"All reputation gained and lost will be restored."}`},
		{69, "reason.deleted.user", `{"name":"deleted","description":"Delete profile, authentication associations."}`},
		{70, "reason.suspended", `{"name":"suspended","description":"A suspended user can’t log in."}`},
		{71, "reason.inactive", `{"name":"inactive","description":"An inactive user must re-validate their email."}`},
		{72, "reason.looks_ok", `{"name":"looks ok","description":"This post is good as-is and not low quality."}`},
		{73, "reason.needs_edit", `{"name":"needs edit, and I did it","description":"Improve and correct problems with this post yourself."}`},
		{74, "reason.needs_close", `{"name":"needs close","description":"A closed question can’t answer, but still can edit, vote and comment."}`},
		{75, "reason.needs_delete", `{"name":"needs delete","description":"All reputation gained and lost will be restored."}`},
		{76, "question.flag.reasons", `["reason.spam","reason.rude_or_abusive","reason.something","reason.a_duplicate"]`},
		{77, "answer.flag.reasons", `["reason.spam","reason.rude_or_abusive","reason.something","reason.not_a_answer"]`},
		{78, "comment.flag.reasons", `["reason.spam","reason.rude_or_abusive","reason.something","reason.no_longer_needed"]`},
		{79, "question.close.reasons", `["reason.a_duplicate","reason.community_specific","reason.not_clarity","reason.something"]`},
		{80, "question.status.reasons", `["reason.normal","reason.closed","reason.deleted"]`},
		{81, "answer.status.reasons", `["reason.normal","reason.deleted"]`},
		{82, "comment.status.reasons", `["reason.normal","reason.deleted"]`},
		{83, "user.status.reasons", `["reason.normal.user","reason.suspended","reason.deleted.user","reason.inactive"]`},
		{84, "question.review.reasons", `["reason.looks_ok","reason.needs_edit","reason.needs_close","reason.needs_delete"]`},
		{85, "answer.review.reasons", `["reason.looks_ok","reason.needs_edit","reason.needs_delete"]`},
		{86, "comment.review.reasons", `["reason.looks_ok","reason.needs_edit","reason.needs_delete"]`},
	}
	_, err := engine.Insert(defaultConfigTable)
	return err
}

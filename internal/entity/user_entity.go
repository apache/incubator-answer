package entity

import "time"

const (
	UserStatusAvailable = 1
	UserStatusSuspended = 9
	UserStatusDeleted   = 10
)

const (
	EmailStatusAvailable    = 1
	EmailStatusToBeVerified = 2
)

const (
	UserAdminFlag = 1
)

// User user
type User struct {
	ID             string    `xorm:"not null pk autoincr comment('user id') BIGINT(20) id"`
	CreatedAt      time.Time `xorm:"created comment('create time') TIMESTAMP created_at"`
	UpdatedAt      time.Time `xorm:"updated comment('update time') TIMESTAMP updated_at"`
	SuspendedAt    time.Time `xorm:"comment('suspended time') TIMESTAMP suspended_at"`
	DeletedAt      time.Time `xorm:"comment('delete time') TIMESTAMP deleted_at"`
	LastLoginDate  time.Time `xorm:"comment('last_login_date') TIMESTAMP last_login_date"`
	Username       string    `xorm:"not null default '' comment('username') VARCHAR(50) username"`
	Pass           string    `xorm:"not null default '' comment('password') VARCHAR(255) pass"`
	EMail          string    `xorm:"not null comment('email') VARCHAR(100) e_mail"`
	MailStatus     int       `xorm:"not null default 2 comment('mail status(1 pass 2 to be verified)') TINYINT(4) mail_status"`
	NoticeStatus   int       `xorm:"not null default 2 comment('notice status(1 on 2off)') INT(11) notice_status"`
	FollowCount    int       `xorm:"not null default 0 comment('follow count') INT(11) follow_count"`
	AnswerCount    int       `xorm:"not null default 0 comment('answer count') INT(11) answer_count"`
	QuestionCount  int       `xorm:"not null default 0 comment('question count') INT(11) question_count"`
	Rank           int       `xorm:"not null default 0 comment('rank') INT(11) rank"`
	Status         int       `xorm:"not null default 1 comment('user status(available: 1; deleted: 10)') INT(11) status"`
	AuthorityGroup int       `xorm:"not null default 1 comment('authority group') INT(11) authority_group"`
	DisplayName    string    `xorm:"not null default '' comment('display name') VARCHAR(50) display_name"`
	Avatar         string    `xorm:"not null default '' comment('avatar') VARCHAR(255) avatar"`
	Mobile         string    `xorm:"not null comment('mobile') VARCHAR(20) mobile"`
	Bio            string    `xorm:"not null comment('bio markdown') TEXT bio"`
	BioHtml        string    `xorm:"not null comment('bio html') TEXT bio_html"`
	Website        string    `xorm:"not null default '' comment('website') VARCHAR(255) website"`
	Location       string    `xorm:"not null default '' comment('location') VARCHAR(100) location"`
	IPInfo         string    `xorm:"not null default '' comment('ip info') VARCHAR(255) ip_info"`
	IsAdmin        bool      `xorm:"not null default 0 comment('admin flag') INT(11) is_admin"`
}

// TableName user table name
func (User) TableName() string {
	return "user"
}

type UserSearch struct {
	User
	Page     int `json:"page" form:"page"`           //Query number of pages
	PageSize int `json:"page_size" form:"page_size"` //Search page size
}

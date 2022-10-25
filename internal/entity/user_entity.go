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
	ID             string    `xorm:"not null pk autoincr BIGINT(20) id"`
	CreatedAt      time.Time `xorm:"created TIMESTAMP created_at"`
	UpdatedAt      time.Time `xorm:"updated TIMESTAMP updated_at"`
	SuspendedAt    time.Time `xorm:"TIMESTAMP suspended_at"`
	DeletedAt      time.Time `xorm:"TIMESTAMP deleted_at"`
	LastLoginDate  time.Time `xorm:"TIMESTAMP last_login_date"`
	Username       string    `xorm:"not null default '' VARCHAR(50) UNIQUE username"`
	Pass           string    `xorm:"not null default '' VARCHAR(255) pass"`
	EMail          string    `xorm:"not null VARCHAR(100) e_mail"`
	MailStatus     int       `xorm:"not null default 2 TINYINT(4) mail_status"`
	NoticeStatus   int       `xorm:"not null default 2 INT(11) notice_status"`
	FollowCount    int       `xorm:"not null default 0 INT(11) follow_count"`
	AnswerCount    int       `xorm:"not null default 0 INT(11) answer_count"`
	QuestionCount  int       `xorm:"not null default 0 INT(11) question_count"`
	Rank           int       `xorm:"not null default 0 INT(11) rank"`
	Status         int       `xorm:"not null default 1 INT(11) status"`
	AuthorityGroup int       `xorm:"not null default 1 INT(11) authority_group"`
	DisplayName    string    `xorm:"not null default '' VARCHAR(30) display_name"`
	Avatar         string    `xorm:"not null default '' VARCHAR(255) avatar"`
	Mobile         string    `xorm:"not null VARCHAR(20) mobile"`
	Bio            string    `xorm:"not null TEXT bio"`
	BioHtml        string    `xorm:"not null TEXT bio_html"`
	Website        string    `xorm:"not null default '' VARCHAR(255) website"`
	Location       string    `xorm:"not null default '' VARCHAR(100) location"`
	IPInfo         string    `xorm:"not null default '' VARCHAR(255) ip_info"`
	IsAdmin        bool      `xorm:"not null default false BOOL is_admin"`
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

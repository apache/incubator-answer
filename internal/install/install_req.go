package install

import (
	"fmt"
	"strings"

	"xorm.io/xorm/schemas"
)

// CheckConfigFileResp check config file if exist or not response
type CheckConfigFileResp struct {
	ConfigFileExist bool `json:"config_file_exist"`
	DbTableExist    bool `json:"db_table_exist"`
}

// CheckDatabaseReq check database
type CheckDatabaseReq struct {
	DbType     string `json:"db_type"`
	DbUsername string `json:"db_username"`
	DbPassword string `json:"db_password"`
	DbHost     string `json:"db_host"`
	DbName     string `json:"db_name"`
	DbFile     string `json:"db_file"`
}

// GetConnection get connection string
func (r *CheckDatabaseReq) GetConnection() string {
	if r.DbType == string(schemas.SQLITE) {
		return r.DbFile
	}
	if r.DbType == string(schemas.MYSQL) {
		return fmt.Sprintf("%s:%s@tcp(%s)/%s",
			r.DbUsername, r.DbPassword, r.DbHost, r.DbName)
	}
	if r.DbType == string(schemas.POSTGRES) {
		host, port := parsePgSQLHostPort(r.DbHost)
		return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
			host, port, r.DbUsername, r.DbPassword, r.DbName)
	}
	return ""
}

func parsePgSQLHostPort(dbHost string) (host string, port string) {
	if strings.Contains(dbHost, ":") {
		idx := strings.LastIndex(dbHost, ":")
		host, port = dbHost[:idx], dbHost[idx+1:]
	} else if len(dbHost) > 0 {
		host = dbHost
	}
	if host == "" {
		host = "127.0.0.1"
	}
	if port == "" {
		port = "5432"
	}
	return host, port
}

// CheckDatabaseResp check database response
type CheckDatabaseResp struct {
	ConnectionSuccess bool `json:"connection_success"`
}

// InitEnvironmentResp init environment response
type InitEnvironmentResp struct {
	Success            bool   `json:"success"`
	CreateConfigFailed bool   `json:"create_config_failed"`
	DefaultConfig      string `json:"default_config"`
	ErrType            string `json:"err_type"`
}

// InitBaseInfoReq init base info request
type InitBaseInfoReq struct {
	Language      string `json:"language"`
	SiteName      string `json:"site_name"`
	SiteURL       string `json:"site_url"`
	ContactEmail  string `json:"contact_email"`
	AdminName     string `json:"admin_name"`
	AdminPassword string `json:"admin_password"`
	AdminEmail    string `json:"admin_email"`
}

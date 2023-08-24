package constant

const (
	UserNormal    = "normal"
	UserSuspended = "suspended"
	UserDeleted   = "deleted"
	UserInactive  = "inactive"
)
const (
	EmailStatusAvailable    = 1
	EmailStatusToBeVerified = 2
)

func ConvertUserStatus(status, mailStatus int) string {
	switch status {
	case 1:
		if mailStatus == EmailStatusToBeVerified {
			return UserInactive
		}
		return UserNormal
	case 9:
		return UserSuspended
	case 10:
		return UserDeleted
	}
	return UserNormal
}

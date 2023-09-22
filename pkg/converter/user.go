package converter

import "github.com/segmentfault/pacman/utils"

func DeleteUserDisplay(userID string) string {
	return utils.EnShortID(StringToInt64(userID), 100)
}

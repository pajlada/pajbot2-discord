package channels

var (
	validChannelRoles = map[string]bool{
		"action-log":        true,
		"moderation-action": true,
		"weeb-channel":      true,
		"system-messages":   true,
	}
)

func ValidRole(channelRole string) (ok bool) {
	_, ok = validChannelRoles[channelRole]
	return
}

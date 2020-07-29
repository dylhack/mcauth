package bot

// SyncHandler is for syncing roles of users in memory to prevent Discord API rate limiting.
type SyncHandler struct {
	sync map[string]*[]string
}

// GetSyncHandler is used by StartBot
func GetSyncHandler() SyncHandler {
	return SyncHandler{
		make(map[string]*[]string),
	}
}

// SyncRoles stores a given member's roles in memory.
func (sh *SyncHandler) SyncRoles(userID string, roles *[]string) {
	sh.sync[userID] = roles
}

// GetRoles returns a member's roles from memory (possibly nil).
func (sh *SyncHandler) GetRoles(userID string) (*[]string, bool) {
	roles, isOK := sh.sync[userID]

	return roles, isOK
}

// GetDiscordIDs returns all the ID's stored.
func (sh *SyncHandler) GetDiscordIDs() (userIDs []string) {
	for userID := range sh.sync {
		userIDs = append(userIDs, userID)
	}
	return userIDs
}

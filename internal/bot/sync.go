package bot

type SyncHandler struct {
	sync map[string]*[]string
}

func GetSyncHandler() SyncHandler {
	return SyncHandler{
		make(map[string]*[]string),
	}
}

func (sh *SyncHandler) SyncRoles(userID string, roles *[]string) {
	sh.sync[userID] = roles
}

func (sh *SyncHandler) GetRoles(userID string) (*[]string, bool) {
	roles, isOK := sh.sync[userID]

	return roles, isOK
}

func (sh *SyncHandler) GetDiscordIDs() (userIDs []string) {
	for userID := range sh.sync {
		userIDs = append(userIDs, userID)
	}
	return userIDs
}

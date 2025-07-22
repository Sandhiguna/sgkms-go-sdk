package authentication

// Buat struct untuk menampung response JSON
type LoginRes struct {
	Result struct {
		SessionToken      string `json:"sessionToken"`
		AgentDN           string `json:"agentDN"`
		LastUsedTime      int64  `json:"lastUsedTime"`
		IdleTimeoutInMins int    `json:"idleTimeoutInMins"`
		ExpiredAt         int64  `json:"expiredAt"`
	} `json:"result"`
}

type RefreshSessionReq struct {
	SlotID       int    `json:"slotId"`
	SessionToken string `json:"sessionToken"`
}

type RandomResult struct {
	Random []byte `json:"random"`
}

type RngResponse struct {
	Result RandomResult `json:"result"`
}
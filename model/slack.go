package model

type Slack struct {
	Command     string `json:"command"`
	User        string `json:"user"`
	Text        string `json:"text"`
	ResponseURL string `json:"response_url"`
	TeamDomain  string `json:"team_domain"`
	ChannelID   string `json:"channel_id"`
	ChannelName string `json:"channel_name"`
	UserID      string `json:"user_id"`
	Token       string `json:"token"`
	TeamID      string `json:"team_id"`
}

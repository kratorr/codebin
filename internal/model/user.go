package model

type User struct {
	ID               int64  `json:"id"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	Username         string `json:"username"`
	Avatar           string `json:"avatar"`
	GoogleID         string `json:"google_id"`
	GithubID         string `json:"github_id"`
	TwitchID         string `json:"twitch_id"`
	ActivationStatus int    `json:"activation_status"`
}

package models

type PebbleLoginResponse struct {
	ClientSecret string `json:"Client-Secret"`
	UID          string `json:"UID"`
}

type UnpackedAccessToken struct {
	Acr   string `json:"acr"`
	UID   string `json:"uid"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type UserSessionRequest struct {
	SessionKey string `json:"session_key"`
	SessionID  string `json:"session_id"`
}

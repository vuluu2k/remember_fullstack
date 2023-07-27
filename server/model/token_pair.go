package model

type TokenPair struct {
	TokenID      string `json:"token_id"`
	RefreshToken string `json:"refresh_token"`
}

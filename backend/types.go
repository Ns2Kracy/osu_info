package main

import "time"

type AccessTokenResponse struct {
	TokenType   string `json:"token_type"`
	ExpiresIn   int32  `json:"expires_in"`
	AccessToken string `json:"access_token"`
	Time        time.Time
}

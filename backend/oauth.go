package main

import (
	"golang.org/x/oauth2"
)

var OsuEndPoint = oauth2.Endpoint{
	AuthURL:  "https://osu.ppy.sh/oauth/authorize",
	TokenURL: "https://osu.ppy.sh/oauth/token",
}

const (
	ClientID     = "14131"
	ClientSecret = "HjMD0JFmHjeWdxspdF6f6H34RllMoikUZQgZn7Em"
	RedirectURL  = "http://localhost:5700/oauth2"
	ResponseType = "code"
	GrantType    = "authorization_code"
)

var conf = &oauth2.Config{
	ClientID:     ClientID,
	ClientSecret: ClientSecret,
	Endpoint:     OsuEndPoint,
	RedirectURL:  RedirectURL,
}

func NewOauth2Client() {

}

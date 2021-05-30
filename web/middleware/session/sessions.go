package session

import "github.com/kataras/iris/v12/sessions"

var (
	cookieNameForSessionID = "mycookiesessionnameid"
)

type Sesstion struct {
	// Sess *sessions.Sessions
}

func DefaultSession() *sessions.Sessions {
	sess := sessions.New(sessions.Config{Cookie: cookieNameForSessionID})
	return sess
}

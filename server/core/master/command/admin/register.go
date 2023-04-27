package admin

import (
	"server/core/master/session/sessions"
	"strings"
)

func Register() {
}

func retroGradient(text string, blacklist string, charColor string, session *sessions.Session) string {
	return strings.ReplaceAll(text, blacklist, charColor+blacklist)
}

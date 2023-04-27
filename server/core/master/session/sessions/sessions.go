package sessions

import (
	"fmt"
	"server/core/master/database"
	"server/core/master/session/sshSession"
	"time"
)

var (
	sessions = make(map[uint64]*Session)
)

type Session struct {
	ID uint64
	*sshSession.SecureShellSession
	*database.User
	Created time.Time
}

func Get(name string) *Session {
	for _, s := range sessions {
		if s.User.Name == name {
			return s
		}
	}
	return nil
}

func Count() int {
	return len(sessions)
}

func Clone() []Session {
	var list []Session
	for _, session := range sessions {
		list = append(list, *session)
	}
	return list
}

func (session *Session) Print(a ...interface{}) error {
	return session.Write([]byte(fmt.Sprint(a...)))
}

func (session *Session) Printf(format string, val ...any) error {
	return session.Write([]byte(fmt.Sprintf(format, val)))
}

func (session *Session) Println(a ...interface{}) error {
	return session.Write([]byte(fmt.Sprint(a...) + "\r\n"))
}

func (session *Session) Clear() error {
	return session.Write([]byte("\033c"))
}

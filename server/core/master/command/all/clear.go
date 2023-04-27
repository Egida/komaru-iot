package all

import (
	"server/core/master/command"
	"server/core/master/session/sessions"
)

func init() {
	command.Make(&command.Command{
		Aliases:     []string{"clear", "cls", "wipe"},
		Description: "Clears the screen",
		Admin:       false,
		Executor: func(args []string, session *sessions.Session) {
			session.Clear()
		},
	})
}

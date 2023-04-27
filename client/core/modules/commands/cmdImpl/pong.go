package cmdImpl

import (
	"client/core/config"
	"client/core/modules/commands"
)

func init() {
	commands.Make(&commands.Command{
		Id: 0,
		Executor: func(client config.Client) error {
			client.Write("\x02")
			return nil
		},
	})
}

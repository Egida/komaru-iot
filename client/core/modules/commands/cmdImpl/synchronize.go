package cmdImpl

import (
	"client/core/config"
	"client/core/modules/commands"
	"fmt"
)

func init() {
	commands.Make(&commands.Command{
		Id: 2,
		Executor: func(client config.Client) error {
			fmt.Println("(cmdImpl/synchronize.go) Received synchronize request")
			err := client.Connection.ReadObject(&config.Bot)
			if err != nil {
				return err
			}
			client.Write("\x01")
			fmt.Println("(cmdImpl/synchronize.go) Successfully synchronized configuration")
			return nil
		},
	})
}

package cmdImpl

import (
	"errors"
	"server/core/config"
	"server/core/slaves/commands"
	"server/core/slaves/slave"
)

func init() {
	commands.Make(&commands.Command{
		Id: 2,
		Executor: func(client *slave.Slave) error {
			err := client.Connection.WriteObject(config.Slave)
			if err != nil {
				return errors.New("failed to send synchronize request")
			}
			return nil
		},
	})
}

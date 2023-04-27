package commands

import "client/core/config"

var (
	commands = make([]*Command, 0)
)

type Command struct {
	Id       int
	Executor func(client config.Client) error
}

func Make(command *Command) {
	commands = append(commands, command)
}

func Clone() []*Command {
	return commands
}

func Get(id int) *Command {
	for _, command := range commands {
		if command.Id == id {
			return command
		}
	}
	return nil
}

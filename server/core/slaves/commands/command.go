package commands

import (
	"errors"
	"fmt"
	"server/core/slaves/slave"
)

var (
	commands = make([]*Command, 0)
)

type Command struct {
	Id       int
	Executor func(client *slave.Slave) error
}

func Make(command *Command) {
	commands = append(commands, command)
}

func Clone() []*Command {
	return commands
}

func Execute(id int, client *slave.Slave) error {
	cmd := Get(id)
	if cmd == nil {
		return errors.New("command does not exist")
	}

	client.Write(fmt.Sprint(id))
	err := cmd.Executor(client)
	if err != nil {
		return err
	}

	return nil
}

func Get(id int) *Command {
	for _, command := range commands {
		if command.Id == id {
			return command
		}
	}
	return nil
}

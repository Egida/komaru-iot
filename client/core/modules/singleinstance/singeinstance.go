package singleinstance

import (
	"client/core/config"
	"fmt"
	"net"
)

// l8r lol
func SingleInstance() bool {
	_, err := net.Dial("tcp", "127.0.0.1:"+fmt.Sprint(config.SINGLEINSTANCE))
	if err == nil {
		return true
	} else {
		go func() {
			for {
				listener, err := net.Listen("tcp", "127.0.0.1:"+fmt.Sprint(config.SINGLEINSTANCE))
				if err != nil {
					continue
				}
				defer listener.Close()
				for {
					conn, err := listener.Accept()
					if err != nil {
						panic(err)
					}
					defer conn.Close()
				}
			}
		}()
		return false
	}
}

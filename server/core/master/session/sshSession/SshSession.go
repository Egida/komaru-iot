package sshSession

import (
	"log"
	"net"

	"golang.org/x/crypto/ssh"
)

type SecureShellSession struct {
	Channel    ssh.Channel
	Conn       net.Conn
	ServerConn ssh.ServerConn
}

func (session SecureShellSession) Write(data []byte) error {
	_, err := session.Channel.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func (session SecureShellSession) Read(data []byte) error {
	_, err := session.Channel.Read(data)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (session SecureShellSession) Close() error {
	err := session.Channel.Close()
	if err != nil {
		return err
	}
	return nil
}

func (session SecureShellSession) CloseWrite() error {
	err := session.Channel.CloseWrite()
	if err != nil {
		return err
	}
	return nil
}

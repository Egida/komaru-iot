package networking

import (
	"fmt"
	"net"
)

func Dial(address string, blowfishKey, chachaKey []byte) (*SecureConnection, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}

	secureConnection := &SecureConnection{
		Connection:  conn,
		BlowfishKey: blowfishKey,
		ChachaKey:   chachaKey,
	}

	err = handleClient(secureConnection)
	if err != nil {
		return secureConnection, err
	}

	return secureConnection, nil
}

func handleClient(client *SecureConnection) error {
	fmt.Println("(networking/handshake.go) Starting Handshake")
	err := client.Write([]byte("\x01"))
	if err != nil {
		return err
	}

	fmt.Println("(networking/handshake.go) Handshake request has been accepted from the server. (0/2)")

	err = client.Write([]byte("\x02"))
	if err != nil {
		return err
	}

	blowfishKey, err := client.Read(1024)
	if err != nil {
		return err
	}

	client.BlowfishKey = blowfishKey

	fmt.Println("(networking/handshake.go) Received new Blowfish key (1/2)")

	err = client.Write([]byte("\x04"))
	if err != nil {
		return err
	}

	chachaKey, err := client.Read(1024)
	if err != nil {
		return err
	}

	client.ChachaKey = chachaKey

	fmt.Println("(networking/handshake.go) Received new ChaCha20 key (2/2)")

	fmt.Println("(networking/handshake.go) Successfully handshaked.")
	return nil
}

package networking

import (
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
	err := client.Write([]byte("\x01"))
	if err != nil {
		return err
	}

	err = client.Write([]byte("\x02"))
	if err != nil {
		return err
	}

	blowfishKey, err := client.Read(1024)
	if err != nil {
		return err
	}

	client.BlowfishKey = blowfishKey

	err = client.Write([]byte("\x04"))
	if err != nil {
		return err
	}

	chachaKey, err := client.Read(1024)
	if err != nil {
		return err
	}

	client.ChachaKey = chachaKey
	return nil
}

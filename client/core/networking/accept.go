package networking

import (
	"net"
)

type Listener struct {
	BlowfishKey,
	ChachaKey []byte
	listener net.Listener
}

func (listener *Listener) Accept() (*SecureConnection, error) {
	conn, err := listener.listener.Accept()
	if err != nil {
		return nil, err
	}

	connection := &SecureConnection{
		Connection:  conn,
		BlowfishKey: listener.BlowfishKey,
		ChachaKey:   listener.ChachaKey,
	}

	var stage = make([]byte, 1024)

	// Stage 1: Verify that the client is the client, so we can go into the actual encryption stage
	stage, err = connection.Read(16)
	if err != nil {
		return connection, err
	}

	if stage[0] != 1 {
		return connection, err
	}

	// Stage 2: Generate a new Blowfish Key & send it.
	stage, err = connection.Read(16)
	if err != nil {
		return connection, err
	}

	if stage[0] != 2 {
		return connection, err
	}

	blowfishKey, err := connection.writeKey(16)
	if err != nil {
		return nil, err
	}

	connection.BlowfishKey = blowfishKey

	// Stage 4: Generate a new ChaCha20 Key & send it.
	stage, err = connection.Read(16)
	if err != nil {
		return connection, err
	}

	if stage[0] != 4 {
		return connection, err
	}

	chachaKey, err := connection.writeKey(32)
	if err != nil {
		return nil, err
	}

	connection.ChachaKey = chachaKey

	return connection, nil
}

func (conn *SecureConnection) writeKey(size int) ([]byte, error) {
	key, err := generateKey(size)
	if err != nil {
		return make([]byte, 0), err
	}

	err = conn.Write(key)
	if err != nil {
		return make([]byte, 0), err
	}

	return key, nil
}

func (listener *Listener) Close() error {
	return listener.listener.Close()
}

func (listener *Listener) Addr() net.Addr {
	return listener.listener.Addr()
}

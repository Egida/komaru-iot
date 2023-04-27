package networking

import (
	"bytes"
	"client/core/networking/encryption"
	"encoding/gob"
	"net"
	"strings"
)

type SecureConnection struct {
	Connection net.Conn
	BlowfishKey,
	TwofishKey,
	ChachaKey []byte
}

func (client *SecureConnection) Write(data []byte) error {
	encrypt, err := encryption.Encrypt(client.BlowfishKey, client.ChachaKey, data)
	if err != nil {
		return err
	}

	_, err = client.Connection.Write(encrypt)
	if err != nil {
		return err
	}
	return nil
}

func (client *SecureConnection) WriteObject(data interface{}) error {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	if err := enc.Encode(data); err != nil {
		return err
	}

	return client.Write(buf.Bytes())
}

func (client *SecureConnection) ReadObject(out interface{}) error {
	input, err := client.Read(8192 * 4)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(input)
	dec := gob.NewDecoder(buf)

	if err := dec.Decode(out); err != nil {
		return err
	}

	return nil
}
func (client *SecureConnection) Read(bufferSize int) (buffer []byte, err error) {
	var realBuffer = make([]byte, bufferSize)

	if _, err = client.Connection.Read(realBuffer); err != nil {
		return realBuffer, err
	}

	decrypted, err := encryption.Decrypt(client.BlowfishKey, client.ChachaKey, []byte(strings.TrimSpace(strings.Replace(string(realBuffer), "\x00", "", -1))))
	if err != nil {
		return realBuffer, err
	}

	return decrypted, nil
}

func (client *SecureConnection) Close() error {
	return client.Connection.Close()
}

func (client *SecureConnection) RemoteAddr() net.Addr {
	return client.Connection.RemoteAddr()
}

func (client *SecureConnection) LocalAddr() net.Addr {
	return client.Connection.LocalAddr()
}

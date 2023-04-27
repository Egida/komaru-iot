package slave

import (
	"server/core/slaves/networking"
)

type Slave struct {
	ID         int
	Connection *networking.SecureConnection
}

func Make(connection *networking.SecureConnection) *Slave {
	return &Slave{
		ID:         -1,
		Connection: connection,
	}
}

func (slave *Slave) Write(data string) {
	slave.Connection.Write([]byte(data))
}

func (slave *Slave) Read(size int) (string, error) {
	read, err := slave.Connection.Read(size)
	if err != nil {
		return "", err
	}
	return string(read), nil
}

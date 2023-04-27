package networking

import "net"

func Listen(address string, BlowfishKey, ChachaKey []byte) (*Listener, error) {
	originalListener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	return &Listener{
		BlowfishKey: BlowfishKey,
		ChachaKey:   ChachaKey,
		listener:    originalListener,
	}, nil
}

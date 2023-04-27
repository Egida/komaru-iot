package sessions

import (
	"crypto/rand"
	"encoding/binary"
)

func New(session *Session) *Session {
	buf := make([]byte, 8)
	for {
		if _, err := rand.Read(buf); err != nil {
			return nil
		}

		id := binary.BigEndian.Uint64(buf)
		if _, ok := sessions[id]; !ok {
			session.ID = id
			sessions[id] = session
			return session
		}
	}
}

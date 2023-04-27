package main

import (
	"client/core/client"
)

var (
	Key = []byte("\x32\x4E\x4A\x64\x31\x4A\x46\x35\x65\x69\x44\x50\x50\x34\x4F\x6D")
)

func main() {
	/*args := os.Args
	if len(args) == 2 {
		client.Serve()
	} else {
		cmd := exec.Command(os.Args[0], "s")
		cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}
		err := cmd.Start()
		if err != nil {
			fmt.Println(err)
			// handle error
		}
		fmt.Println("komaru <3 :heart:")
	}*/
	client.Serve()
}

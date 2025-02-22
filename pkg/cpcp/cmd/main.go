package main

import (
	"fmt"

	"github.com/harnyk/commie/pkg/cpcp"
)

func main() {
	host := cpcp.NewProcessClient("/bin/bash")

	if err := host.Start(); err != nil {
		panic(err)
	}

	go func() {
		for line := range host.Receive() {
			fmt.Println("Received:", line)
		}
	}()

	host.Send("echo Hello")
	host.Send("exit")

	go func() {
		for err := range host.Errors() {
			fmt.Println("Error:", err)
		}
	}()

	exitCode := <-host.ExitCode()
	fmt.Println("Plugin exited with code:", exitCode)

	host.Stop()
}

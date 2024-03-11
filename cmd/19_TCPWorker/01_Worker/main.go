package main

import (
	"fmt"
	"net"
	"os"
)

func init() {
	runCmd("bash", "cleanpng.sh")
}

func main() {
	listener, errListen := net.Listen("tcp", socket)
	if errListen != nil {
		fmt.Println("listener error: ", errListen)

		os.Exit(1)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("failed to accept conn:", err)

			continue
		}

		handleConn(conn) // could be also concurrent.
	}
}

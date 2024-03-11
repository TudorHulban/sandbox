package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func handleConn(connection net.Conn) {
	defer connection.Close()

	reader := bufio.NewReader(connection)
	for {
		bytes, errReader := reader.ReadBytes(byte('\n'))
		if errReader != nil {
			if errReader != io.EOF {
				fmt.Println("failed to read data", errReader)
			}

			return
		}

		request := string(bytes)
		fmt.Println("request: ", request)

		resp, err := doWork(request)
		if err != nil {
			connection.Write([]byte("try later" + "\n"))

			continue
		}

		connection.Write([]byte(resp + "\n"))
	}
}

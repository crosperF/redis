package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	// buffer_test()
	// return

	lisner, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := lisner.Accept()

	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	for {
		buffer := make([]byte, 1024)

		n, err := conn.Read(buffer)
		fmt.Println(n)
		fmt.Println(n)
		if err != nil {
			fmt.Println(err)

			if err == io.EOF {
				fmt.Println("Client has closed the connection. EOF reached.")
				break
			}
			fmt.Println("error reading from client: ", err.Error())
			os.Exit(1)
		}
		conn.Write([]byte("+OK\r\n"))
	}

}

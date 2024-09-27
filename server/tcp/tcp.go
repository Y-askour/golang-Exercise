package tcp

import (
	"fmt"
	"net"
)

func Listener() {
	Listener, error := net.Listen("tcp", ":8080")
	defer Listener.Close()
	if error != nil {
		fmt.Println("creating a listener failed")
	}
	for true {
		connection, error := Listener.Accept()
		if error != nil {
			fmt.Println("accepting a connection failed")
		}
		go handleConnection(&connection)
	}
}

func handleConnection(connection *net.Conn) {
	defer (*connection).Close()
	buf := make([]byte, 1024)

	for {
		len, err := (*connection).Read(buf)
		if err != nil {
			break
		}

		message := string(buf[:len])

		fmt.Printf(message)
	}
}

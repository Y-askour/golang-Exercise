package tcp

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
)

type RandomMessage struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

func Listener() {
	Listener, error := net.Listen("tcp", ":8080")
	defer Listener.Close()
	if error != nil {
		log.Fatalf("listen")
	}
	for true {
		connection, error := Listener.Accept()
		if error != nil {
			log.Fatalf("accept")
		}
		go handleConnection(&connection)
	}
}

func handleConnection(connection *net.Conn) {
	defer (*connection).Close()
	buf := make([]byte, 1024)
	var message_as_struct RandomMessage

	for {
		len, error := (*connection).Read(buf)
		if error != nil {
			if error == io.EOF {
				break
			}
			return
		}

		// encode
		json.Unmarshal(buf[:len], &message_as_struct)
		message_as_a_binary, err := json.Marshal(message_as_struct)
		if err != nil {
			log.Fatalf("marshal")
		}

		// print fields of the struct
		fmt.Println(message_as_struct.Count)
		fmt.Println(message_as_struct.Name)

		//send the message to telegraf
		outgoingConn, error := net.Dial("tcp", "localhost:8094")
		defer outgoingConn.Close()
		if error != nil {
			log.Fatalf("dial")
		}
		outgoingConn.Write(message_as_a_binary)

	}
}

package tcp

import (
	"encoding/json"
	"log"
	"net"
	"time"
)

// my json struct
type RandomMessage struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

var count = 1

const HOST = "server"
const PORT = "8080"

// try connect goroutine  and then read from the channel of the message goroutine
func connectAndSend(channel chan []byte) {
	message := <-channel

	// connect to server
	conn, err := net.Dial("tcp", HOST+":"+PORT)
	if err != nil {
		log.Fatalf("Dial")
	}
	defer conn.Close()

	// send data
	data := []byte(message)
	_, err = conn.Write(data)
	if err != nil {
		log.Fatalf("write")
	}
}

func MainLoop() {
	channel := make(chan []byte)
	for {
		go connectAndSend(channel)

		// message  goroutine -> create message and encode it
		message_as_a_struct := RandomMessage{"younes", count}
		message_as_a_binary, err := json.Marshal(message_as_a_struct)

		if err != nil {
			log.Fatalf("marshal")
		}
		channel <- message_as_a_binary
		count++
		time.Sleep(time.Second)
		//
	}
}

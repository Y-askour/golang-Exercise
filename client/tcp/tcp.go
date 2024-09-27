package tcp

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

var count = 1

const HOST = "localhost"
const PORT = "8080"

// try connect goroutine  and then read from the channel of the message goroutine
func connectAndSend(channel chan string) {
	message := <-channel

	// connect to server
	conn, err := net.Dial("tcp", HOST+":"+PORT)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	// send data
	data := []byte(message)
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println(err)
	}
}

func MainLoop() {
	channel := make(chan string)
	for {
		go connectAndSend(channel)

		// message goroutine -> create message
		time.Sleep(time.Second)
		channel <- "younes " + strconv.Itoa(count)
		count++
		//
	}
}

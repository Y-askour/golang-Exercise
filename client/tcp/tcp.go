package tcp

import (
	"fmt"
	"strconv"
	"time"
)

var count = 1

// try connect goroutine  and then read from the channel of the message goroutine
func connect(channel chan string) {
	message := <-channel
	fmt.Println(message)
}

func MainLoop() {
	channel := make(chan string)
	for {
		go connect(channel)
		// message goroutine -> create message
		time.Sleep(time.Second * 2)
		channel <- "younes " + strconv.Itoa(count)
		count++
	}
}

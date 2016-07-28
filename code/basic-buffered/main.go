package main

import "fmt"

func main() {
	ping := make(chan string, 2)

	ping <- "buffer"
	ping <- "channel"
	// ping <- "channel 2"

	fmt.Println(<-ping)
	fmt.Println(<-ping)
}

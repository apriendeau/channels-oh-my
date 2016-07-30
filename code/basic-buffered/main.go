package main

import "fmt"

func main() {
	ping := make(chan string, 2)

	ping <- "buffer"
	ping <- "channel"

	// if I was to uncomment this line:
	// things would go horribly bad and block till
	// it was unblocked, in this case: DEADLOCK
	// ping <- "channel 2"

	fmt.Println(<-ping)
	fmt.Println(<-ping)
}

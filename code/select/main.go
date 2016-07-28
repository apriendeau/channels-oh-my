package main

import "fmt"

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() { ch2 <- "landon" }()
	go func() { ch1 <- "hayden" }()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Printf("Channel 1: %s\n", msg)
		case msg := <-ch2:
			fmt.Printf("Channel 2: %s\n", msg)
		}
	}
}

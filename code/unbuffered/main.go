package main

import "fmt"

func main() {
	msg := make(chan string)

	go func() {
		for i := 0; i <= 100; i++ {
			msg <- "hello"
		}
	}()

	messages := <-msg
	fmt.Println(messages)
}

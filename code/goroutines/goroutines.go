package main

import "fmt"

func cook(food string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("Cooking %s:%d\n", food, i)
	}
}

func main() {
	cook("chicken")
	go cook("cake")

	fmt.Println("done")
}

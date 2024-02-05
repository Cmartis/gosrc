package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	//get greeting message and print it
	message := greetings.Hello("Chris")
	fmt.Println(message)
}

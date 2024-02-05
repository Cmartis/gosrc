package greetings

import "fmt"

// hello returns greetings for a named person
func Hello(name string) string {
	//return a greeting that embedds the namne of person
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

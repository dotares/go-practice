package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hello, %v. Welcome to Go", name)
	return message 
}
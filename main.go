package main

import (
	"fmt"
	 "github.com/lwrabbit/greetings/v2"
)

// Hi returns a friendly greeting
func Hi(name string) string {
	return fmt.Sprintf("%s, %s", greetings.HI, name)
}

// Hello returns a friendly greeting
func Hello(name string) string {
	return fmt.Sprintf("%s, %s", greetings.HELLO, name)
}

func main() {
	fmt.Println(Hi("lwrabbit"))
	fmt.Println(Hello("lwrabbit"))
}
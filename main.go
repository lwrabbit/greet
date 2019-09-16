package main

import (
	"fmt"
	 "github.com/lwrabbit/greetings"
)

// Hi returns a friendly greeting
func Hi(name string) string {
	return fmt.Sprintf("%s, %s", greetings.HI, name)
}

func main() {
	fmt.Println(Hi("lwrabbit"))
}
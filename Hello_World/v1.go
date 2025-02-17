package main

import (
	"fmt"
	"strings"
)

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if len(strings.TrimSpace(name)) == 0 {
		return englishHelloPrefix + "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Ongchen"))
}

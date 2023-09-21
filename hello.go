package main

import "fmt"

const englishprefixname = "Hello, "

func hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishprefixname + name
}

func main() {
	fmt.Println(hello(""))
}

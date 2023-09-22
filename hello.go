package main

import "fmt"

const spanish = "Spanish"
const englishprefixname = "Hello, "
const spanishprefixname = "Hola, "

func hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	if lang == spanish {
		return spanishprefixname + name
	}

	return englishprefixname + name
}

func main() {
	fmt.Println(hello("", "Spanish"))
}

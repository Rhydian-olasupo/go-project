package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishprefixname = "Hello, "
const spanishprefixname = "Hola, "
const frenchprefixname = "Bonjour, "

func hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishprefixname

	switch lang {
	case french:
		prefix = frenchprefixname
	case spanish:
		prefix = spanishprefixname
	}

	return prefix + name
}

func main() {
	fmt.Println(hello("Rhydian", "French"))
}

package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"

	englishprefixname = "Hello, "
	spanishprefixname = "Hola, "
	frenchprefixname  = "Bonjour, "
)

func hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case french:
		prefix = frenchprefixname
	case spanish:
		prefix = spanishprefixname
	default:
		prefix = englishprefixname
	}

	return prefix
}

func main() {
	fmt.Println(hello("", "French"))
}

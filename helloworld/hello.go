package main

import "fmt"

const (
	french   = "French"
	japanese = "Japanese"
	spanish  = "Spanish"

	englishHelloPrefix  = "Hello, "
	frenchHelloPrefix   = "Bonjour, "
	japaneseHelloPrefix = "こんにちは, "
	spanishHelloPrefix  = "Hola, "
)

func main() {
	fmt.Println(Hello("world", ""))
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case japanese:
		prefix = japaneseHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

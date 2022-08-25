package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, lang string) string {

	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch lang {
	case "french":
		prefix = frenchHelloPrefix
	case "spanish":
		prefix = spanishHelloPrefix
	}
	return prefix + name

}

func main() {
	fmt.Println(Hello("World", "English"))
}

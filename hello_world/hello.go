package main

const spanish = "Spanish"
const french = "French"
const world = "World"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

// Hello returns a hello-world string
func Hello(name string, language string) string {
	return greetingPrefix(language) + greetingName(name)
}

func greetingPrefix(language string) string {
	switch language {
	case spanish:
		return spanishHelloPrefix
	case french:
		return frenchHelloPrefix
	default:
		return englishHelloPrefix
	}
}

func greetingName(name string) string {
	if name == "" {
		name = "World"
	}

	return name
}

func main() {}

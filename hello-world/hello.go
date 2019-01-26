package main

const helloPrefix = "Hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "french":
		prefix = "frenchHelloPrefix"
	case "spanish":
		prefix = "spanishHelloPrefix"
	default:
		prefix = "englishPrefix"
	}

	return
}

func main() {
}

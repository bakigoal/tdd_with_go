package hello

import "fmt"

func Hello(name, language string) string {
	return fmt.Sprintf("%s, %s!", greeting(language), name)
}

var greetings = map[string]string{
	"ru": "Привет",
	"es": "Hola",
	"fr": "Bonjour",
}

const defaultGreeting = "Hello"

func greeting(language string) string {
	greeting, exists := greetings[language]
	if exists {
		return greeting
	}
	return defaultGreeting
}

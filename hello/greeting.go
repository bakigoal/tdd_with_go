package hello

import "fmt"

func Hello(name, language string) string {
	greet := greeting(language)
	if name == "" {
		name = greet.defaultName
	}
	return fmt.Sprintf("%s, %s!", greet.greeting, name)
}

type Greet struct {
	greeting    string
	defaultName string
}

var greetings = map[string]Greet{
	"ru": {"Привет", "мир"},
	"es": {"Hola", "Espania"},
	"fr": {"Bonjour", "Paris"},
}

var defaultGreeting = Greet{"Hello", "world"}

func greeting(language string) Greet {
	greeting, exists := greetings[language]
	if exists {
		return greeting
	}
	return defaultGreeting
}

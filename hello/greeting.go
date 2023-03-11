package hello

import "fmt"

func Hello(name, language string) string {
	if name == "" {
		name = getDefaultName(language)
	}
	return fmt.Sprintf("%s, %s!", greeting(language), name)
}

var prefixesByLanguage = map[string]string{
	"ru": "Привет",
	"es": "Hola",
	"fr": "Bonjour",
}

var defaultNamesByLanguage = map[string]string{
	"ru": "Мир",
}

const (
	defaultGreeting = "Hello"
	defaultName     = "World"
)

func greeting(language string) string {
	greeting, exists := prefixesByLanguage[language]
	if exists {
		return greeting
	}
	return defaultGreeting
}

func getDefaultName(language string) string {
	name, exists := defaultNamesByLanguage[language]
	if exists {
		return name
	}
	return defaultName
}

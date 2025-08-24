package airportrobot

import "fmt"

// Greeter interface
type Greeter interface {
	// LanguageName which returns the name of the language (a string) that the robot is supposed to greet the visitor in.
	LanguageName() string
	// Greet which accepts a visitor's name (a string) and returns a string with the greeting message in a specific language.
	Greet(visitorName string) string
}

func SayHello(visitorName string, greeter Greeter) string {
	return fmt.Sprintf("I can speak %v: %v!", greeter.LanguageName(), greeter.Greet(visitorName))
}

type German struct{}

func (German) LanguageName() string {
	return "German"
}

func (German) Greet(visitorName string) string {
	return fmt.Sprintf("Hallo %s", visitorName)
}

type Italian struct{}

func (Italian) LanguageName() string {
	return "Italian"
}

func (Italian) Greet(visitorName string) string {
	return fmt.Sprintf("Ciao %s", visitorName)
}

type Portuguese struct{}

func (Portuguese) LanguageName() string {
	return "Portuguese"
}

func (Portuguese) Greet(visitorName string) string {
	return fmt.Sprintf("Olá %s", visitorName)
}

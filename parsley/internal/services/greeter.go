package services

import "fmt"

type Greeter interface {
	SayHello(name string, polite bool) string
}

type greeter struct {
	salutation string
}

// SayHello Generates a greeter message for the given user.
func (g *greeter) SayHello(name string, polite bool) string {
	if polite {
		return fmt.Sprintf("Good day, %s!\n", name)
	} else {
		return fmt.Sprintf("%s, %s\n", g.salutation, name)
	}
}

// NewGreeterFactory The activator function for Greeter services. Returns a constructor function that resolves Greeter instances.
func NewGreeterFactory(salutation string) func() Greeter {
	return func() Greeter {
		return &greeter{salutation: salutation}
	}
}

package services

import "fmt"

type Greeter interface {
	SayHello(name string, polite bool) string
}

type greeter struct{}

// SayHello Generates a greeter message for the given user.
func (g *greeter) SayHello(name string, polite bool) string {
	if polite {
		return fmt.Sprintf("Good day, %s!\n", name)
	} else {
		return fmt.Sprintf("Hi, %s\n", name)
	}
}

// NewGreeterFactory The activator function for Greeter service instances.
func NewGreeterFactory() Greeter {
	return &greeter{}
}

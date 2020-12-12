// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package wire_lib

import (
	"fmt"
)

// Injectors from wire_lib.go:

func InitializeEvent() Event {
	message := NewMessage()
	greeter := NewGreeter(message)
	event := NewEvent(greeter)
	return event
}

// wire_lib.go:

// Message
type Message string

// Greeter
type Greeter struct {
	Message Message // 信息
}

// Event
type Event struct {
	Greeter Greeter
}

// NewMessage
func NewMessage() Message {
	return Message("Hi wire")
}

// NewGreeter
func NewGreeter(msg Message) Greeter {
	return Greeter{Message: msg}
}

// Greet
func (g *Greeter) Greet() Message {
	return g.Message
}

// NewEvent
func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

// Start
func (e *Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

// Run
func Run() {
	event := InitializeEvent()
	event.Start()
}

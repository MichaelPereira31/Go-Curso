package events

import "time"

type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
}

type EventHandlerInterface interface {
	Handle(event EventInterface) error
}



type EventDispatcherInterface interface {
	Register(eventNames string, handle EventHandlerInterface) error
	Dispatcher(event EventInterface) error
	Remove(eventNames string, handle EventHandlerInterface) error
	Has(eventNames string) bool
}
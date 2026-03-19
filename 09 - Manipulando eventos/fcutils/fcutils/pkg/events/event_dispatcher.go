package events

import "errors"

var msgEventAlreadyRegistered = "handler already registered for another event"
type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (ed *EventDispatcher) Register(eventNames string, handle EventHandlerInterface) error {
	if _, ok := ed.handlers[eventNames]; !ok {
		for _, h := range ed.handlers[eventNames] {
			if h == handle {
				return errors.New(msgEventAlreadyRegistered)
			}
		}
	}
	ed.handlers[eventNames] = append(ed.handlers[eventNames], handle)
	return nil
}
package events

import (
	"sync"
	"time"
)

// Event
type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
	SetPayload(payload interface{})
}

// Event operation
type EventHandlerInterface interface {
	Handle(event EventInterface, wg *sync.WaitGroup)
}

// Event manager: register and dispatch events
type EventDispatcherInterface interface {
	Register(eventName string, handler EventHandlerInterface) error
	Dispatch(event EventInterface) error
	Remove(eventName string, handler EventHandlerInterface) error
	Has(eventName string, handler EventHandlerInterface) bool
	Clear()
}

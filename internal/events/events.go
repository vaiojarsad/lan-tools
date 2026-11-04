package events

import (
	"errors"
)

type Handler[T any] func(T) error
type InMemoryBus[T any] struct {
	handlers []Handler[T]
}

func NewInMemoryBus[T any]() *InMemoryBus[T] {
	return &InMemoryBus[T]{}
}

func (b *InMemoryBus[T]) Subscribe(handler Handler[T]) {
	b.handlers = append(b.handlers, handler)
}

func (b *InMemoryBus[T]) Publish(event T) error {
	var errs []error
	for _, handler := range b.handlers {
		if err := handler(event); err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}

type IPChanged struct {
	ISPCode string
	OldIP   string
	NewIP   string
}

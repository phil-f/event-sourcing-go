package event

import (
	"log"
)

type Handler interface {
	Handle(event *Event) error
}

type EventBus interface {
	Publish(event *Event)
	Subscribe(handler Handler)
}

type Bus struct {
	subscriptions []*subscription
}

func NewBus() *Bus {
	return &Bus{subscriptions: []*subscription{}}
}

func (b *Bus) Publish(event *Event) {
	for _, s := range b.subscriptions {
		s.ch <- event
	}
}

func (b *Bus) Subscribe(handler Handler) {
	s := newSubscription(handler.Handle)
	b.subscriptions = append(b.subscriptions, s)
	s.start()
}

type subscription struct {
	ch     chan *Event
	handle func(*Event) error
}

func newSubscription(handle func(*Event) error) *subscription {
	return &subscription{
		ch:     make(chan *Event, 100),
		handle: handle,
	}
}

func (s *subscription) start() {
	go func(s *subscription) {
		for {
			if err := s.handle(<-s.ch); err != nil {
				log.Println(err)
			}
		}
	}(s)
}

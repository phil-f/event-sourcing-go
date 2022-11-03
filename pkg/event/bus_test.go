package event

import (
	"testing"
	"time"
)

func TestBus(t *testing.T) {
	bus := NewBus()
	handler := &handlerStub{}
	bus.Subscribe(handler)
	event := &Event{0, "hello"}
	bus.Publish(event)

	time.Sleep(100 * time.Millisecond)

	if handler.counter != 1 {
		t.Fatalf("bad counter")
	}
	if event != handler.event {
		t.Fatalf("bad event")
	}
}

type handlerStub struct {
	event   *Event
	counter int
}

func (h *handlerStub) Handle(event *Event) error {
	h.event = event
	h.counter++
	return nil
}

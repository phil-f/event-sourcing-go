package event

import (
	"reflect"
	"spt-go/pkg/uuid"
	"testing"
)

func TestStore_Save(t *testing.T) {
	bus := &busStub{}
	store := NewStore(bus)
	id := uuid.New()
	events := []*Event{
		{0, "something"},
		{1, "thing"},
	}

	if err := store.Save(id, events, 0); err != nil {
		t.Fatalf(err.Error())
	}
	if !reflect.DeepEqual(bus.ReceivedEvents, events) {
		t.Fatalf("whoops")
	}
	if bus.Called != 2 {
		t.Fatalf("whoopsie")
	}
}

type busStub struct {
	ReceivedEvents []*Event
	Called         int
}

func (b *busStub) Publish(event *Event) {
	b.ReceivedEvents = append(b.ReceivedEvents, event)
	b.Called++
}

func (b *busStub) Subscribe(handler Handler) {}

package aggregate

import (
	"spt-go/pkg/event"
	"spt-go/pkg/uuid"
	"testing"
)

func TestAggregate_ApplyChange(t *testing.T) {
	r := &handlerStub{}
	event := &event.Event{Data: "something"}
	r.ApplyChange(r, event)

	changes := r.GetUncommittedChanges()
	if len(changes) != 1 {
		t.Fatalf("whoops")
	}

	eventsEqual := event == r.GetUncommittedChanges()[0]
	if !eventsEqual {
		t.Fatalf("whoops")
	}
}

func TestAggregate_MarkChangesAsCommitted(t *testing.T) {
	r := &handlerStub{}
	event := &event.Event{Data: "something"}

	r.ApplyChange(r, event)
	r.MarkChangesAsCommitted()

	changes := r.GetUncommittedChanges()
	if len(changes) != 0 {
		t.Fatalf("whoops")
	}
}

func TestAggregate_Rehydrate(t *testing.T) {
	r := &handlerStub{}
	events := []*event.Event{
		{Data: "something"},
		{Data: "thing"},
	}
	r.Rehydrate(r, events)

	changes := r.GetUncommittedChanges()
	if len(changes) != 0 {
		t.Fatalf("whoops")
	}
}

func TestAggregate_GetID(t *testing.T) {
	r := &handlerStub{}
	r.Version = 3
	if r.GetVersion() != r.Version {
		t.Fatalf("whoops")
	}
}

func TestAggregate_GetVersion(t *testing.T) {
	r := &handlerStub{}
	r.ID = uuid.New()
	if r.GetID() != r.ID {
		t.Fatalf("whoops")
	}
}

type handlerStub struct {
	Aggregate
}

func (r *handlerStub) Apply(*event.Event) {}

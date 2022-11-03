package aggregate

import (
	"spt-go/pkg/event"
	"spt-go/pkg/uuid"
)

type Handler interface {
	Root
	Apply(*event.Event)
}

type Root interface {
	GetUncommittedChanges() []*event.Event
	MarkChangesAsCommitted()
	Rehydrate(root Handler, events []*event.Event)
	GetVersion() int
	GetID() uuid.UUID
}

type Aggregate struct {
	uncommittedEvents []*event.Event
	ID                uuid.UUID
	Version           int
}

func (a *Aggregate) GetUncommittedChanges() []*event.Event {
	if a.uncommittedEvents == nil {
		return []*event.Event{}
	}

	return a.uncommittedEvents
}

func (a *Aggregate) MarkChangesAsCommitted() {
	a.uncommittedEvents = []*event.Event{}
}

func (a *Aggregate) Rehydrate(handler Handler, events []*event.Event) {
	for _, e := range events {
		a.applyChange(false, handler, e)
	}
}

func (a *Aggregate) GetVersion() int {
	return a.Version
}

func (a *Aggregate) GetID() uuid.UUID {
	return a.ID
}

func (a *Aggregate) ApplyChange(handler Handler, event *event.Event) {
	a.applyChange(true, handler, event)
}

func (a *Aggregate) applyChange(isNew bool, handler Handler, event *event.Event) {
	event.Version = a.Version
	handler.Apply(event)
	if isNew {
		a.uncommittedEvents = append(a.uncommittedEvents, event)
	}
}

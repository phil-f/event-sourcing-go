package event

import (
	"fmt"
	"reflect"
	"spt-go/pkg/uuid"
)

type EventStore interface {
	Save(aggregateID uuid.UUID, events []*Event, expectedVersion int) error
	GetEvents(aggregateID uuid.UUID) ([]*Event, error)
}

type Record struct {
	ID      uuid.UUID
	Type    string
	Event   *Event
	Version int
}

func newEventRecord(id uuid.UUID, evt *Event, version int) *Record {
	return &Record{
		ID:      id,
		Type:    reflect.TypeOf(evt).String(),
		Event:   evt,
		Version: version,
	}
}

type Store struct {
	bus    EventBus
	events map[uuid.UUID][]*Record
}

func NewStore(bus EventBus) *Store {
	return &Store{
		bus:    bus,
		events: make(map[uuid.UUID][]*Record),
	}
}

func (s *Store) Save(aggregateID uuid.UUID, events []*Event, expectedVersion int) error {
	eventRecords, ok := s.events[aggregateID]
	if ok && len(eventRecords) > 0 {
		currentVersion := eventRecords[len(eventRecords)-1].Version
		if currentVersion != expectedVersion && expectedVersion != -1 {
			return fmt.Errorf("unexpected version, expected: %v but current is: %v", expectedVersion, currentVersion)
		}
	}

	version := expectedVersion
	for _, e := range events {
		version++
		e.Version = version
		eventRecord := newEventRecord(aggregateID, e, version)
		s.events[aggregateID] = append(s.events[aggregateID], eventRecord)
		s.bus.Publish(eventRecord.Event)
	}
	return nil
}

func (s *Store) GetEvents(aggregateID uuid.UUID) ([]*Event, error) {
	recs, ok := s.events[aggregateID]
	if ok {
		var events []*Event
		for _, r := range recs {
			events = append(events, &Event{Data: r.Event.Data})
		}
		return events, nil
	}
	return nil, fmt.Errorf("no events found for aggregate with id %s", aggregateID)
}

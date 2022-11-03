package event

type Event struct {
	Version int
	Data    any
}

func NewEvent(eventData any) *Event {
	return &Event{Data: eventData}
}

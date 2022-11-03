package event

import (
	"testing"
)

func TestNewEvent(t *testing.T) {
	data := "hello"
	event := NewEvent(data)

	if event.Data != data {
		t.Fatalf("whoops")
	}
	if event.Version != 0 {
		t.Fatalf("whoopsie")
	}
}

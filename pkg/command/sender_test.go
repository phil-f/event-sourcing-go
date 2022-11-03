package command

import (
	"reflect"
	"testing"
)

func TestSender_Send(t *testing.T) {
	command := &commandStub{Data: "hello"}
	commandType := reflect.TypeOf(command)
	handler := &handlerStub{}
	handlers := map[reflect.Type]Handler{
		commandType: handler,
	}
	factory := NewFactory(handlers)
	sender := NewSender(factory)

	if err := sender.Send(command); err != nil {
		t.Fatalf(err.Error())
	}
	if handler.ReceivedCommand != command {
		t.Fatalf("oh no")
	}
}

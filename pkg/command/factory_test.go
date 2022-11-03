package command

import (
	"reflect"
	"testing"
)

func TestHandlerFactory_GetHandler(t *testing.T) {
	command := commandStub{}
	commandType := reflect.TypeOf(command)
	handlers := map[reflect.Type]Handler{
		commandType: &handlerStub{},
	}
	factory := NewFactory(handlers)

	handler, err := factory.GetHandler(reflect.TypeOf(command))
	if err != nil {
		t.Fatalf("whoops")
	}
	if handler != handler {
		t.Fatalf("bad handler")
	}
}

type commandStub struct {
	Data any
}

type handlerStub struct {
	ReceivedCommand any
	Called          int
}

func (h *handlerStub) Handle(cmd any) error {
	h.ReceivedCommand = cmd
	h.Called++
	return nil
}

package command

import (
	"errors"
	"fmt"
	"reflect"
)

type Handler interface {
	Handle(cmd any) error
}

type CommandHandlerFactory interface {
	GetHandler(t reflect.Type) (Handler, error)
}

type HandlerFactory struct {
	handlersByType map[reflect.Type]Handler
}

func NewFactory(handlersByType map[reflect.Type]Handler) *HandlerFactory {
	return &HandlerFactory{handlersByType: handlersByType}
}

func (f *HandlerFactory) GetHandler(t reflect.Type) (Handler, error) {
	h, ok := f.handlersByType[t]
	if !ok {
		return nil, errors.New(fmt.Sprintf("Unable to find handler for command of type: %s", t))
	}
	return h, nil
}

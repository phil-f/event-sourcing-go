package command

import (
	"reflect"
)

type CommandSender interface {
	Send(cmd any) error
}

type Sender struct {
	handlerFactory CommandHandlerFactory
}

func NewSender(factory CommandHandlerFactory) *Sender {
	return &Sender{handlerFactory: factory}
}

func (r *Sender) Send(cmd any) error {
	h, err := r.handlerFactory.GetHandler(reflect.TypeOf(cmd))
	if err != nil {
		return err
	}
	if err := h.Handle(cmd); err != nil {
		return err
	}
	return nil
}

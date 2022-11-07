package inventory

import (
	"errors"
	"fmt"
	"spt-go/pkg/aggregate"
	"spt-go/pkg/event"
	"spt-go/pkg/uuid"
)

type ItemPreview struct {
	ID   uuid.UUID
	Name string
}

type Item struct {
	aggregate.Aggregate
	activated bool
}

func (i *Item) Apply(event *event.Event) {
	switch data := event.Data; data.(type) {
	case ItemCreated:
		i.ID = data.(ItemCreated).ID
		i.activated = true
	case ItemDeactivated:
		i.activated = false
	}
}

func (i *Item) Create(id uuid.UUID, name string) {
	i.ApplyChange(i, event.NewEvent(ItemCreated{id, name}))
}

func (i *Item) Deactivate() error {
	if !i.activated {
		return fmt.Errorf("unable to deactive item with id: %v as it is already deactivated", i.ID)
	}
	i.ApplyChange(i, event.NewEvent(ItemDeactivated{i.ID}))
	return nil
}

func (i *Item) Remove(count int) error {
	if count <= 0 {
		return fmt.Errorf("unable to remove %v items as the value is not above 0", count)
	}
	i.ApplyChange(i, event.NewEvent(ItemsRemovedFromInventory{i.ID, count}))
	return nil
}

func (i *Item) CheckIn(count int) error {
	if count <= 0 {
		return fmt.Errorf("unable to check in %v items as the value is not above 0", count)
	}
	i.ApplyChange(i, event.NewEvent(ItemsCheckedInToInventory{i.ID, count}))
	return nil
}

func (i *Item) Rename(newName string) error {
	if newName == "" {
		return errors.New("can't rename an inventory item to an empty string")
	}
	i.ApplyChange(i, event.NewEvent(ItemRenamed{i.ID, newName}))
	return nil
}

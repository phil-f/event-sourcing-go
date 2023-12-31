package inventory

import (
	"event-sourcing-go/pkg/aggregate"
	"event-sourcing-go/pkg/event"
	"event-sourcing-go/pkg/uuid"
)

type InventoryRepository interface {
	GetByID(id uuid.UUID) (*Item, error)
	Save(a aggregate.Root, expectedVersion int) error
}

type Repository struct {
	es event.EventStore
}

func NewRepository(es event.EventStore) *Repository {
	return &Repository{es: es}
}

func (r *Repository) Save(a aggregate.Root, expectedVersion int) error {
	if err := r.es.Save(a.GetID(), a.GetUncommittedChanges(), expectedVersion); err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetByID(id uuid.UUID) (*Item, error) {
	inventoryItem := &Item{}

	events, err := r.es.GetEvents(id)
	if err != nil {
		return nil, err
	}
	inventoryItem.Rehydrate(inventoryItem, events)
	return inventoryItem, nil
}

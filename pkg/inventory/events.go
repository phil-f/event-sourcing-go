package inventory

import (
	"event-sourcing-go/pkg/uuid"
)

type ItemDeactivated struct {
	ID uuid.UUID
}

type ItemCreated struct {
	ID   uuid.UUID
	Name string
}

type ItemRenamed struct {
	ID   uuid.UUID
	Name string
}

type ItemsCheckedInToInventory struct {
	ID    uuid.UUID
	Count int
}

type ItemsRemovedFromInventory struct {
	ID    uuid.UUID
	Count int
}

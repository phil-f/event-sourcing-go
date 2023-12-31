package inventory

import (
	"event-sourcing-go/pkg/uuid"
)

type DeactivateInventoryItem struct {
	ID      uuid.UUID
	Version int
}

type CreateInventoryItem struct {
	ID   uuid.UUID
	Name string
}

type RenameInventoryItem struct {
	ID      uuid.UUID
	NewName string
	Version int
}

type CheckInItemsToInventory struct {
	ID      uuid.UUID
	Count   int
	Version int
}

type RemoveItemsFromInventory struct {
	ID      uuid.UUID
	Count   int
	Version int
}

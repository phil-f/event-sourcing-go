package inventory

import (
	"spt-go/pkg/uuid"
)

type ItemModel struct {
	ID      uuid.UUID
	Name    string
	Count   int
	Version int
}

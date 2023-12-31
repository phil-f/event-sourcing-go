package inventory

import (
	"event-sourcing-go/pkg/uuid"
)

type itemRecord struct {
	ID           uuid.UUID
	Name         string
	CurrentCount int
	Version      int
}

type InventoryDatabase interface {
	UpsertInventoryItemRecord(rec *itemRecord)
	GetInventoryItemRecord(id uuid.UUID) *itemRecord
	GetInventoryItemRecords() []*itemRecord
	RemoveInventoryItemRecord(id uuid.UUID)
}

type Database struct {
	records map[uuid.UUID]*itemRecord
}

func NewDatabase() *Database {
	return &Database{records: make(map[uuid.UUID]*itemRecord)}
}

func (d *Database) UpsertInventoryItemRecord(rec *itemRecord) {
	d.records[rec.ID] = rec
}

func (d *Database) GetInventoryItemRecord(id uuid.UUID) *itemRecord {
	return d.records[id]
}

func (d *Database) GetInventoryItemRecords() []*itemRecord {
	items := make([]*itemRecord, 0, len(d.records))
	for _, v := range d.records {
		items = append(items, v)
	}
	return items
}

func (d *Database) RemoveInventoryItemRecord(id uuid.UUID) {
	delete(d.records, id)
}

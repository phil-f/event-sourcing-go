package inventory

import (
	"event-sourcing-go/pkg/event"
	"fmt"
)

type Projector struct {
	db InventoryDatabase
}

func NewProjector(db InventoryDatabase) *Projector {
	return &Projector{db: db}
}

func (p *Projector) Handle(event *event.Event) error {
	switch t := event.Data.(type) {
	case ItemCreated:
		p.createInventoryItem(event.Data.(ItemCreated))
	case ItemRenamed:
		p.renameInventoryItem(event.Data.(ItemRenamed), event.Version)
	case ItemsRemovedFromInventory:
		p.removeInventoryItem(event.Data.(ItemsRemovedFromInventory), event.Version)
	case ItemsCheckedInToInventory:
		p.checkInInventoryItem(event.Data.(ItemsCheckedInToInventory), event.Version)
	case ItemDeactivated:
		p.deactivateInventoryItem(event.Data.(ItemDeactivated))
	default:
		return fmt.Errorf("unable to find handler for event of type: %T", t)
	}
	return nil
}

func (p *Projector) createInventoryItem(event ItemCreated) {
	item := &itemRecord{
		ID:           event.ID,
		Name:         event.Name,
		CurrentCount: 0,
		Version:      0,
	}

	p.db.UpsertInventoryItemRecord(item)
}

func (p *Projector) renameInventoryItem(event ItemRenamed, version int) {
	record := p.db.GetInventoryItemRecord(event.ID)
	record.Name = event.Name
	record.Version = version
	p.db.UpsertInventoryItemRecord(record)
}

func (p *Projector) removeInventoryItem(event ItemsRemovedFromInventory, version int) {
	record := p.db.GetInventoryItemRecord(event.ID)
	record.CurrentCount -= event.Count
	record.Version = version
	p.db.UpsertInventoryItemRecord(record)
}

func (p *Projector) checkInInventoryItem(event ItemsCheckedInToInventory, version int) {
	record := p.db.GetInventoryItemRecord(event.ID)
	record.CurrentCount += event.Count
	record.Version = version
	p.db.UpsertInventoryItemRecord(record)
}

func (p *Projector) deactivateInventoryItem(event ItemDeactivated) {
	p.db.RemoveInventoryItemRecord(event.ID)
}

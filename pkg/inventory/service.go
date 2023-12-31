package inventory

import (
	"event-sourcing-go/pkg/command"
	"event-sourcing-go/pkg/uuid"
)

type InventoryItemService interface {
	CreateInventoryItem(id uuid.UUID, name string) error
	RenameInventoryItem(id uuid.UUID, name string, version int) error
	RemoveInventoryItems(id uuid.UUID, count, version int) error
	CheckInInventoryItems(id uuid.UUID, count, version int) error
	DeactivateInventoryItem(id uuid.UUID, version int) error
	GetInventoryItemPreviews() []*ItemPreview
	GetInventoryItemDetails(id uuid.UUID) *ItemModel
}

type ItemService struct {
	commandSender     command.CommandSender
	inventoryDatabase InventoryDatabase
}

func NewItemService(sender command.CommandSender, database InventoryDatabase) *ItemService {
	return &ItemService{
		commandSender:     sender,
		inventoryDatabase: database,
	}
}

func (s *ItemService) CreateInventoryItem(id uuid.UUID, name string) error {
	cmd := CreateInventoryItem{
		ID:   id,
		Name: name,
	}
	return s.commandSender.Send(cmd)
}

func (s *ItemService) RenameInventoryItem(id uuid.UUID, name string, version int) error {
	cmd := RenameInventoryItem{
		ID:      id,
		NewName: name,
		Version: version,
	}
	return s.commandSender.Send(cmd)
}

func (s *ItemService) RemoveInventoryItems(id uuid.UUID, count, version int) error {
	cmd := RemoveItemsFromInventory{
		ID:      id,
		Count:   count,
		Version: version,
	}
	return s.commandSender.Send(cmd)
}

func (s *ItemService) CheckInInventoryItems(id uuid.UUID, count, version int) error {
	cmd := CheckInItemsToInventory{
		ID:      id,
		Count:   count,
		Version: version,
	}
	return s.commandSender.Send(cmd)
}

func (s *ItemService) DeactivateInventoryItem(id uuid.UUID, version int) error {
	cmd := DeactivateInventoryItem{
		ID:      id,
		Version: version,
	}
	return s.commandSender.Send(cmd)
}

func (s *ItemService) GetInventoryItemPreviews() []*ItemPreview {
	itemRecords := s.inventoryDatabase.GetInventoryItemRecords()
	itemPreviews := make([]*ItemPreview, 0, len(itemRecords))

	for _, v := range itemRecords {
		itemPreviews = append(
			itemPreviews, &ItemPreview{
				ID:   v.ID,
				Name: v.Name,
			},
		)
	}

	return itemPreviews
}

func (s *ItemService) GetInventoryItemDetails(id uuid.UUID) *ItemModel {
	rec := s.inventoryDatabase.GetInventoryItemRecord(id)
	return &ItemModel{
		ID:      rec.ID,
		Name:    rec.Name,
		Count:   rec.CurrentCount,
		Version: rec.Version,
	}
}

package inventory

import (
	"fmt"
)

type Handler struct {
	repo InventoryRepository
}

func NewHandler(repo InventoryRepository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) Handle(cmd any) error {
	switch t := cmd.(type) {
	case CreateInventoryItem:
		return h.createInventoryItem(cmd.(CreateInventoryItem))
	case DeactivateInventoryItem:
		return h.deactivateInventoryItem(cmd.(DeactivateInventoryItem))
	case RemoveItemsFromInventory:
		return h.removeItemsFromInventory(cmd.(RemoveItemsFromInventory))
	case CheckInItemsToInventory:
		return h.checkInItemsToInventory(cmd.(CheckInItemsToInventory))
	case RenameInventoryItem:
		return h.renameInventoryItem(cmd.(RenameInventoryItem))
	default:
		return fmt.Errorf("unable to find a handler for command of type: %s", t)
	}
}

func (h *Handler) createInventoryItem(cmd CreateInventoryItem) error {
	item := &Item{}
	item.Create(cmd.ID, cmd.Name)
	if err := h.repo.Save(item, -1); err != nil {
		return err
	}
	return nil
}

func (h *Handler) deactivateInventoryItem(cmd DeactivateInventoryItem) error {
	item, err := h.repo.GetByID(cmd.ID)
	if err != nil {
		return err
	}
	if err := item.Deactivate(); err != nil {
		return err
	}
	if err := h.repo.Save(item, cmd.Version); err != nil {
		return err
	}
	return nil
}

func (h *Handler) removeItemsFromInventory(cmd RemoveItemsFromInventory) error {
	item, err := h.repo.GetByID(cmd.ID)
	if err != nil {
		return err
	}
	if err := item.Remove(cmd.Count); err != nil {
		return err
	}
	if err := h.repo.Save(item, cmd.Version); err != nil {
		return err
	}
	return nil
}

func (h *Handler) checkInItemsToInventory(cmd CheckInItemsToInventory) error {
	item, err := h.repo.GetByID(cmd.ID)
	if err != nil {
		return err
	}
	if err := item.CheckIn(cmd.Count); err != nil {
		return err
	}
	if err := h.repo.Save(item, cmd.Version); err != nil {
		return err
	}
	return nil
}

func (h *Handler) renameInventoryItem(cmd RenameInventoryItem) error {
	item, err := h.repo.GetByID(cmd.ID)
	if err != nil {
		return err
	}
	if err := item.Rename(cmd.NewName); err != nil {
		return err
	}
	if err := h.repo.Save(item, cmd.Version); err != nil {
		return err
	}
	return nil
}

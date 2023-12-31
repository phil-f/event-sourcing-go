package api

import (
	"event-sourcing-go/pkg/inventory"
	"event-sourcing-go/pkg/templates"
	"event-sourcing-go/pkg/uuid"
	"net/http"
	"path"
	"strconv"
)

type RouteHandler struct {
	templates        *templates.Model
	inventoryItemSvc inventory.InventoryItemService
}

func NewRouteHandler(templates *templates.Model, inventoryItemSvc inventory.InventoryItemService) *RouteHandler {
	return &RouteHandler{
		templates:        templates,
		inventoryItemSvc: inventoryItemSvc,
	}
}

func (h *RouteHandler) Index(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		items := h.inventoryItemSvc.GetInventoryItemPreviews()

		if err := h.templates.Home.Execute(rw, items); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
		return
	}
}

func (h *RouteHandler) AddItem(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if err := h.templates.FormAdd.Execute(rw, nil); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if r.Method == http.MethodPost {
		name := r.FormValue("Name")

		if err := h.inventoryItemSvc.CreateInventoryItem(uuid.New(), name); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}

		redirectToIndex(rw, r)
	}
}

func (h *RouteHandler) Details(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := path.Base(r.URL.Path)
		itemId, _ := uuid.Parse(id)
		item := h.inventoryItemSvc.GetInventoryItemDetails(itemId)

		if err := h.templates.Details.Execute(rw, item); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (h *RouteHandler) Rename(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := path.Base(r.URL.Path)
		itemId, _ := uuid.Parse(id)
		item := h.inventoryItemSvc.GetInventoryItemDetails(itemId)

		if err := h.templates.FormRename.Execute(rw, item); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if r.Method == http.MethodPost {
		id := path.Base(r.URL.Path)
		itemId, _ := uuid.Parse(id)
		name := r.FormValue("Name")
		version, _ := strconv.Atoi(r.FormValue("Version"))

		if err := h.inventoryItemSvc.RenameInventoryItem(itemId, name, version); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}

		redirectToIndex(rw, r)
	}
}

func (h *RouteHandler) CheckIn(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := path.Base(r.URL.Path)
		itemId, _ := uuid.Parse(id)
		item := h.inventoryItemSvc.GetInventoryItemDetails(itemId)

		if err := h.templates.FormCheckIn.Execute(rw, item); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if r.Method == http.MethodPost {
		id := path.Base(r.URL.Path)
		itemId, _ := uuid.Parse(id)
		count, _ := strconv.Atoi(r.FormValue("Number"))
		version, _ := strconv.Atoi(r.FormValue("Version"))

		if err := h.inventoryItemSvc.CheckInInventoryItems(itemId, count, version); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}

		redirectToIndex(rw, r)
	}
}

func (h *RouteHandler) Deactivate(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := path.Base(r.URL.Path)
		itemId, _ := uuid.Parse(id)
		item := h.inventoryItemSvc.GetInventoryItemDetails(itemId)

		if err := h.templates.FormDeactivate.Execute(rw, item); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if r.Method == http.MethodPost {
		id := path.Base(r.URL.Path)
		itemId, _ := uuid.Parse(id)
		version, _ := strconv.Atoi(r.FormValue("Version"))

		if err := h.inventoryItemSvc.DeactivateInventoryItem(itemId, version); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}

		redirectToIndex(rw, r)
	}
}

func (h *RouteHandler) Remove(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := path.Base(r.URL.Path)
		itemId, _ := uuid.Parse(id)
		item := h.inventoryItemSvc.GetInventoryItemDetails(itemId)

		if err := h.templates.FormRemove.Execute(rw, item); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if r.Method == http.MethodPost {
		id := path.Base(r.URL.Path)
		itemId, _ := uuid.Parse(id)
		count, _ := strconv.Atoi(r.FormValue("Number"))
		version, _ := strconv.Atoi(r.FormValue("Version"))

		if err := h.inventoryItemSvc.RemoveInventoryItems(itemId, count, version); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}

		redirectToIndex(rw, r)
	}
}

func redirectToIndex(rw http.ResponseWriter, r *http.Request) {
	http.Redirect(rw, r, "/", http.StatusSeeOther)
}

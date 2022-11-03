package main

import (
	"html/template"
	"net/http"
	"reflect"
	"spt-go/pkg/api"
	"spt-go/pkg/command"
	"spt-go/pkg/event"
	"spt-go/pkg/inventory"
	"spt-go/pkg/templates"
)

func main() {
	eventBus := event.NewBus()
	eventStore := event.NewStore(eventBus)
	inventoryRepo := inventory.NewRepository(eventStore)
	inventoryHandler := inventory.NewHandler(inventoryRepo)
	inventoryDb := inventory.NewDatabase()
	inventoryProjector := inventory.NewProjector(inventoryDb)

	eventBus.Subscribe(inventoryProjector)

	cmdHandlers := map[reflect.Type]command.Handler{
		reflect.TypeOf(inventory.CreateInventoryItem{}):      inventoryHandler,
		reflect.TypeOf(inventory.RenameInventoryItem{}):      inventoryHandler,
		reflect.TypeOf(inventory.DeactivateInventoryItem{}):  inventoryHandler,
		reflect.TypeOf(inventory.CheckInItemsToInventory{}):  inventoryHandler,
		reflect.TypeOf(inventory.RemoveItemsFromInventory{}): inventoryHandler,
	}

	cmdFactory := command.NewFactory(cmdHandlers)
	cmdSender := command.NewSender(cmdFactory)
	inventorySvc := inventory.NewItemService(cmdSender, inventoryDb)
	siteTemplates := setupTemplates()
	routeHandler := api.NewRouteHandler(siteTemplates, inventorySvc)

	_ = serve(routeHandler)
}

func serve(handler *api.RouteHandler) error {
	http.Handle("/web/css/", http.StripPrefix("/web/css/", http.FileServer(http.Dir("web/css"))))

	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/Home/Add", handler.AddItem)
	http.HandleFunc("/Home/CheckIn/", handler.CheckIn)
	http.HandleFunc("/Home/Deactivate/", handler.Deactivate)
	http.HandleFunc("/Home/Details/", handler.Details)
	http.HandleFunc("/Home/Rename/", handler.Rename)
	http.HandleFunc("/Home/Remove/", handler.Remove)

	return http.ListenAndServe(":8080", nil)
}

func setupTemplates() *templates.Model {
	tmpl := template.Must(template.ParseGlob("web/templates/*"))

	return &templates.Model{
		Home:           tmpl.Lookup("home"),
		Details:        tmpl.Lookup("details"),
		FormAdd:        tmpl.Lookup("form-add"),
		FormRename:     tmpl.Lookup("form-rename"),
		FormCheckIn:    tmpl.Lookup("form-check-in"),
		FormDeactivate: tmpl.Lookup("form-deactivate"),
		FormRemove:     tmpl.Lookup("form-remove"),
	}
}

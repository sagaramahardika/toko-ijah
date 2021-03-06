package app

import (
	"fmt"
	"log"
	"net/http"
	
	"toko-ijah/api/app/handler"
	"toko-ijah/api/app/model"
	"toko-ijah/api/config"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)


// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// App initialize with predefined configuration
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		config.DB.Charset)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

// Set all required routers
func (a *App) setRouters() {

	// Routing API CRUD Item
	a.Get("/api/items", a.GetAllItem)
	a.Post("/api/item/create", a.CreateItem)
	a.Get("/api/item/{sku}", a.GetItem)
	a.Put("/api/item/update/{sku}", a.UpdateItem)
	a.Delete("/api/item/{sku}", a.DeleteItem)

	// Routing API CRUD Incoming Item
	a.Get("/api/incoming/items", a.GetAllIncomingItem)
	a.Post("/api/incoming/item/create", a.CreateIncomingItem)
	a.Get("/api/incoming/item/{id}", a.GetIncomingItem)
	a.Put("/api/incoming/item/update/{id}", a.UpdateIncomingItem)
	a.Delete("/api/incoming/item/{id}", a.DeleteIncomingItem)

	// Routing API CRUD Outgoing Item
	a.Get("/api/outgoing/items", a.GetAllOutgoingItem)
	a.Post("/api/outgoing/item/create", a.CreateOutgoingItem)
	a.Get("/api/outgoing/item/{id}", a.GetOutgoingItem)
	a.Put("/api/outgoing/item/update/{id}", a.UpdateOutgoingItem)
	a.Delete("/api/outgoing/item/{id}", a.DeleteOutgoingItem)
}

// Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Wrap the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Wrap the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Handlers to manage Item Data
func (a *App) GetAllItem(w http.ResponseWriter, r *http.Request) {
	handler.GetAllItem(a.DB, w, r)
}

func (a *App) CreateItem(w http.ResponseWriter, r *http.Request) {
	handler.CreateItem(a.DB, w, r)
}

func (a *App) GetItem(w http.ResponseWriter, r *http.Request) {
	handler.GetItem(a.DB, w, r)
}

func (a *App) UpdateItem(w http.ResponseWriter, r *http.Request) {
	handler.UpdateItem(a.DB, w, r)
}

func (a *App) DeleteItem(w http.ResponseWriter, r *http.Request) {
	handler.DeleteItem(a.DB, w, r)
}

// Handlers to manage Incoming Item Data
func (a *App) GetAllIncomingItem(w http.ResponseWriter, r *http.Request) {
	handler.GetAllIncomingItem(a.DB, w, r)
}

func (a *App) CreateIncomingItem(w http.ResponseWriter, r *http.Request) {
	handler.CreateIncomingItem(a.DB, w, r)
}

func (a *App) GetIncomingItem(w http.ResponseWriter, r *http.Request) {
	handler.GetIncomingItem(a.DB, w, r)
}

func (a *App) UpdateIncomingItem(w http.ResponseWriter, r *http.Request) {
	handler.UpdateIncomingItem(a.DB, w, r)
}

func (a *App) DeleteIncomingItem(w http.ResponseWriter, r *http.Request) {
	handler.DeleteIncomingItem(a.DB, w, r)
}

// Handlers to manage Outgoing Item Data
func (a *App) GetAllOutgoingItem(w http.ResponseWriter, r *http.Request) {
	handler.GetAllOutgoingItem(a.DB, w, r)
}

func (a *App) CreateOutgoingItem(w http.ResponseWriter, r *http.Request) {
	handler.CreateOutgoingItem(a.DB, w, r)
}

func (a *App) GetOutgoingItem(w http.ResponseWriter, r *http.Request) {
	handler.GetOutgoingItem(a.DB, w, r)
}

func (a *App) UpdateOutgoingItem(w http.ResponseWriter, r *http.Request) {
	handler.UpdateOutgoingItem(a.DB, w, r)
}

func (a *App) DeleteOutgoingItem(w http.ResponseWriter, r *http.Request) {
	handler.DeleteOutgoingItem(a.DB, w, r)
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"sync"
)

type App struct {
	Router      *mux.Router
	SubRouter   *mux.Router
	TbBukaLapak *gorm.DB
}

func (app *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("GET")
}

func (app *App) Concurrent(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup

	wg.Add(1)
	go Info(&wg,w, r)
	wg.Wait()
}

func (app *App) NoconcuRent(w http.ResponseWriter, r *http.Request) {



	InfoNoconc(w, r)

}

func (app *App) setRouters() {
	app.Get("/concur", app.Concurrent)
	app.Get("/Noconcur", app.NoconcuRent)
}
func (app *App) Initialize() {

	app.Router = mux.NewRouter()

	app.setRouters()
}

func (app *App) Run(host string) {
	headersOk := handlers.AllowedHeaders([]string{"Content-Type", "Access-Control-Allow-Headers", "Authorization", "X-Requested-With", "Access-Control-Allow-Origin", "x-access-token"})
	corsObj := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	log.Fatal(http.ListenAndServe(host, handlers.CORS(corsObj, headersOk, methodsOk)(app.Router)))

}

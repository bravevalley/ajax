package app

import (
	"html/template"
	"log"
	"net/http"
)

type App struct {
	Router *Router
	Server *http.Server
}

func (app *App) Init(tpl template.Template) {
	app.Router = &Router{}
	app.Router.Init(tpl)

}

func (app *App) Start() {
	app.Server = &http.Server{
		Addr: ":8080",
		Handler: app.Router.Mux,
	}

	log.Fatal(app.Server.ListenAndServe())
}
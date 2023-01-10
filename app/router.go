package app

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	Mux *mux.Router
}

func (router *Router) Init(tpl template.Template) {
	router.Mux = mux.NewRouter()
	router.ServeContent(tpl)
}

func (router *Router) ServeContent(tpl template.Template) {
	handlers := &Handler{
		tpl,
	}

	router.Mux.HandleFunc("/signup", handlers.signUp)
	router.Mux.HandleFunc("/signin", handlers.signIn)
	router.Mux.HandleFunc("/dashboard", handlers.dashboard)
	router.Mux.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

}
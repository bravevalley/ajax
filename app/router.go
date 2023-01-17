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

	router.Mux.HandleFunc("/", handlers.dashboard)
	router.Mux.HandleFunc("/signup", handlers.signUp)
	router.Mux.HandleFunc("/login", handlers.signIn)
	router.Mux.HandleFunc("/dashboard", handlers.dashboard)
	router.Mux.HandleFunc("/logout", handlers.logout)
	router.Mux.HandleFunc("/checkuser", handlers.checkuser)
	router.Mux.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

}

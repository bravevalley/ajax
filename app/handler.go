package app

import (
	"html/template"
	"net/http"
)

type Handler struct {
	template template.Template
}

func (handler *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	
}

func (handler *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	
}

func (handler *Handler) dashboard(w http.ResponseWriter, r *http.Request) {
	
}
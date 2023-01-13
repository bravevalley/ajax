package app

import (
	"html/template"
	"log"
	"net/http"

	query "dev.go/databases"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	template template.Template
}

func (handler *Handler) signIn(w http.ResponseWriter, req *http.Request) {	
	c, err := req.Cookie("cbdk");
	if err != nil {

		if req.Method == http.MethodPost {
			us := req.FormValue("username")
			ps := req.FormValue("password")

			ok := query.CheckData(us)
			if !ok {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			passwd, err := query.GetData(us)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				log.Fatalln(err)
			}

			bcrypt.CompareHashAndPassword([]byte(ps), []byte(passwd))

		}
		// No cookie

	}
}

func (handler *Handler) signUp(w http.ResponseWriter, req *http.Request) {

}

func (handler *Handler) dashboard(w http.ResponseWriter, req *http.Request) {

}

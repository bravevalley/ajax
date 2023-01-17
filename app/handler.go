package app

import (
	"html/template"
	"log"
	"net/http"
	"time"

	query "dev.go/databases"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	template template.Template
}

func (handler *Handler) signIn(w http.ResponseWriter, req *http.Request) {
	loadusers()
	c, err := req.Cookie("cbdk")
	if err == nil {
		if ok := query.CheckData(c.Value); ok {
			http.Redirect(w, req, "/dashboard", http.StatusSeeOther)
			return
		}
	}

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

		err = bcrypt.CompareHashAndPassword([]byte(passwd), []byte(ps))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Fatalln(err)
		}

		sid := uuid.Must(uuid.NewRandom())

		c = &http.Cookie{
			Name:  "cbdk",
			Value: sid.String(),
		}

		err = query.SetData(sid.String(), us, 0*time.Second)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Fatalln(err)
		}

		http.SetCookie(w, c)

		http.Redirect(w, req, "/dashboard", http.StatusSeeOther)
		return
	}

	handler.template.ExecuteTemplate(w, "login.gohtml", nil)
}

func (handler *Handler) signUp(w http.ResponseWriter, req *http.Request) {
	loadusers()
	c, err := req.Cookie("cbdk")
	if err == nil {
		if ok := query.CheckData(c.Value); ok {
			http.Redirect(w, req, "/dashboard", http.StatusSeeOther)
			return
		}
	}

	if req.Method == http.MethodPost {
		us := req.FormValue("username")
		ps := req.FormValue("password")
		em := req.FormValue("email")

		passwd, err := bcrypt.GenerateFromPassword([]byte(ps), bcrypt.MinCost)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Fatal(err)
		}

		err = query.InputUser(us,passwd, em)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Fatal(err)
		}

		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	handler.template.ExecuteTemplate(w, "signup.gohtml", nil)
}

func (handler *Handler) dashboard(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("cbdk")
	if err != nil {
		if ok := query.CheckData(c.Value); ok {
			http.Redirect(w, req, "/signin", http.StatusSeeOther)
			return
		}
	}

	if ok := query.CheckData(c.Value); !ok {
		http.Redirect(w, req, "/signin", http.StatusSeeOther)
		return
	}

	username, err := query.GetData(c.Value)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatalln(err)
		return
	}

	handler.template.ExecuteTemplate(w, "dashboard.gohtml", username)
	return

}

func logout(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("cbdk")
	if err != nil {
		if ok := query.CheckData(c.Value); ok {
			http.Redirect(w, req, "/signin", http.StatusSeeOther)
			return
		}
	}

	if err = query.RemoveData(c.Value); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatalln(err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name: "cbdk",
		Value: c.Value,
		MaxAge: -1,
	})

}

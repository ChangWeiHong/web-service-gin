package main

import (
	"html/template"
	"net/http"
)

type ContactInformations struct {
	Email    string
	Nickname string
	Subject  string
	Message  string
}

func forms() {
	template := template.Must(template.ParseFiles("forms.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			template.Execute(w, nil)
			return
		}

		informations := ContactInformations{
			Email:    r.FormValue("email"),
			Nickname: r.FormValue("nickname"),
			Subject:  r.FormValue("subject"),
			Message:  r.FormValue("message"),
		}

		_ = informations

		template.Execute(w, struct{ Success bool }{true})
	})
}

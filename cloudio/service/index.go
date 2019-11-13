package service

import (
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	page := template.Must(template.ParseFiles("assets/index.html"))
	page.Execute(w, nil)
}

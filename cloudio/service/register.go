package service

import (
	"html/template"
	"net/http"
)

func register(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	username := req.Form["username"][0]
	password := req.Form["password"][0]
	mail := req.Form["mail"][0]
	phone := req.Form["phone"][0]

	page := template.Must(template.ParseFiles("assets/info.html"))

	page.Execute(w, map[string]string{
		"username":  username,
		"password":  password,
		"mail": mail,
		"phone": phone,
	})
}

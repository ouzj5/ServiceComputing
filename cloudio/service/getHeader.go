package service

import (
	"net/http"
)

func getHeader(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("register"))
}

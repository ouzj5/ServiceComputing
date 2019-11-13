package service

import (
	"net/http"
)

func unKnown(reqw http.ResponseWriter, req *http.Request) {
	http.Error(reqw, "501 Unkown!", 501)
}

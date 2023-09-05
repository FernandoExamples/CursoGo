package controllers

import "net/http"

func HomeGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

package controllers

import "net/http"

func PingController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

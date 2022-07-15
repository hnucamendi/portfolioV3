package main

import "net/http"

func WebApp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

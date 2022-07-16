package main

import "net/http"

func webApp(wr http.ResponseWriter, r *http.Request) {
	wr.Write([]byte("Hello World"))
}

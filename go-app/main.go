package main

import (
	"log"
	"net/http"
)

var PongHandler = func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func main() {

	http.HandleFunc("/ping", PongHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln("unable to start server", err)
	}
}

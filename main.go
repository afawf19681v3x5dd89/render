package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("asdasd"))
	})

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("текст"))
	})

	err := http.ListenAndServe("0.0.0.0:8000", nil)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

}

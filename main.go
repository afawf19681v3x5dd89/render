package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("asdasd"))
	})

	err := http.ListenAndServe("localhost:80", nil)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

}

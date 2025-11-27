package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("."))

	http.HandleFunc("/toon/example/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Disposition", "inline")
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		fs.ServeHTTP(w, r)
	})

	http.Handle("/", fs)

	log.Println("Serving presentation on http://localhost:8080 ...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

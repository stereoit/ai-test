package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"stereoit.com/my-blank-workspace/phones"
)

func main() {
	r := chi.NewRouter()

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Get("/phones", func(w http.ResponseWriter, r *http.Request) {
		phone := phones.Phone{ID: "1", Name: "iPhone 13", Brand: "Apple"}
		render.JSON(w, r, phone)
	})

	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", r)
}
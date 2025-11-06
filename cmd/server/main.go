package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/go-chi/chi/v5"
)

func main() {
    r := chi.NewRouter()

    r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Server is running 🚀"))
    })

    fmt.Println("Server running on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
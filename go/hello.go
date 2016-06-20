package main

import (
    "fmt"
    "net/http"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hey hey, Skyliner.")
}

func main() {
    http.HandleFunc("/", handler)
    addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
    http.ListenAndServe(addr, nil)
}

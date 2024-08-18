package main

import (
    "flag"
    "fmt"
    "log"
    "net/http"
)

func main() {
    folder := flag.String("d", ".", "folder to publish.")
    port := flag.String("p", "8080", "port to serve.")
    flag.Parse()
    log.Printf("Http static file hosting path \"%s\" on :%s", *folder,  *port)
    file_server := http.FileServer(http.Dir(*folder))
    err := http.ListenAndServe(fmt.Sprintf(":%s", *port), customHeaders(file_server))
    if err != nil {
        log.Fatal(err)
    }
}

func customHeaders(file_server http.Handler) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Cache-Control", "no-cache")
        w.Header().Set("x-server", "Hi from golang-simple-httpserver.")
        file_server.ServeHTTP(w, r)
    }
}

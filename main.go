package main

import (
    "fmt"
    "net/http"
)

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {

        // route will default to / if it doesn't match using net/http default servemux
        // have to make sure to check the / handler if it doesn't match throw a 404 back!
        if r.URL.Path != "/" {
            w.WriteHeader(http.StatusNotFound)
            fmt.Fprintf(w, "%s route not found!", r.URL.Path)
            return
        }

        fmt.Fprintf(w, "Hello World!")
    })

    mux.HandleFunc("/ping", func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "pong!")
    })

    fmt.Println("listing on :3000....")

    http.ListenAndServe(":3000", mux)
}

package main

import (
    "os"
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "skain.tech handler: %s!", r.URL.Path[1:])
}
func containerhandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "skain.tech container ID ", os.Getenv("HOSTNAME"))
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/container", containerhandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

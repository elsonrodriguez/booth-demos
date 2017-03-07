package main

import (
    "fmt"
    "net/http"
    "strings"
    "log"
    "os"
)

func hello(w http.ResponseWriter, r *http.Request) {
    pod_name := os.Getenv("POD_NAME")
    pod_nodename := os.Getenv("POD_NODENAME")
    fmt.Fprintf(w, "Serving you from pod %s on node %s", pod_name, pod_nodename)
}

func main() {
    http.HandleFunc("/", hello) // set router
    listen_port := os.Getenv("LISTEN_PORT")
    listen_address := os.Getenv("LISTEN_ADDRESS")
    listen_spec := strings.Join([]string{listen_address, listen_port}, ":")
    err := http.ListenAndServe(listen_spec, nil) // set listen port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

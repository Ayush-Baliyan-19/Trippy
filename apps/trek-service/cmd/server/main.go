package main

import (
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Trek service ok"))
    })

    log.Println("Trek service running on :8001")
    http.ListenAndServe(":8001", nil)
}

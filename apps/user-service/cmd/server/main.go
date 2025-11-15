package main

import (
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("User -service ok"))
    })

    log.Println("User service running on :8002")
    http.ListenAndServe(":8002", nil)
}

package main

import (
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Match service ok"))
    })

    log.Println("Match service running on :8003")
    http.ListenAndServe(":8003", nil)
}

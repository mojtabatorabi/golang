package main

import (
    //"log"
    "net/http"

    "example.com/go-rest-api/internal/user"
    "example.com/go-rest-api/pkg/logger"
)

func main() {
    // Init logger
    log := logger.New()

    // Setup routes
    http.HandleFunc("/users", user.Handler)

    log.Info("Starting server on :8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Error("Server failed: ", err)
    }
}

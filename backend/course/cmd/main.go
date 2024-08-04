package main

import (
    "course_service/config"
    "course_service/handlers"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    db, err := config.InitDB()
    if err != nil {
        log.Fatalf("Could not set up database: %v", err)
    }

    router := mux.NewRouter()
    handlers.RegisterCourseRoutes(router, db)

    log.Println("Starting course service on port 8081...")
    if err := http.ListenAndServe(":8081", router); err != nil {
        log.Fatalf("Could not start server: %v", err)
    }
}

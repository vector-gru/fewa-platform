package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/lesi/tutor_booking_system/login/handlers"
    "github.com/lesi/tutor_booking_system/login/services"
)

func main() {
    // Initialize auth service with the registration service URL
    authService := services.NewAuthService("http://localhost:8080") // Adjust to your registration service URL

    // Initialize login handler
    loginHandler := handlers.NewLoginHandler(authService)

    // Create router
    router := mux.NewRouter()

    // Serve static files
    router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../static/"))))

    // Define routes
    router.HandleFunc("/api/login", loginHandler.Login).Methods("POST")

    // Serve the login HTML template
    router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "../registration/static/login.html")
    })

    // Start server
    log.Println("Starting login service on port 8082...")
    log.Fatal(http.ListenAndServe(":8082", router))
}

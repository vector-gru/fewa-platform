package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/lesi/tutor_booking_system/registration/pkg/database"
    "github.com/lesi/tutor_booking_system/registration/handlers"
    "github.com/lesi/tutor_booking_system/registration/services"
)

func main() {
    // Initialize database and handle any errors
    db, err := database.InitDB()
    if err != nil {
        log.Fatalf("Error initializing database: %v", err)
    }

    // Create new user service with the database connection
    userService := services.NewUserService(db)
    authService := services.NewAuthService(db) // Initialize AuthService

    // Initialize user handler with the user service and auth service
    userHandler := handlers.NewUserHandler(userService, authService)

    // Create router
    r := mux.NewRouter()

    // Serve static files from the 'static' directory
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../../static/"))))

    // Define routes
    r.HandleFunc("/register", userHandler.RegisterUser).Methods("POST")
    r.HandleFunc("/users", userHandler.GetAllUsers).Methods("GET")
    r.HandleFunc("/users/{id:[0-9]+}", userHandler.GetUserByID).Methods("GET")
    r.HandleFunc("/users/{id:[0-9]+}", userHandler.UpdateUser).Methods("PUT")
    r.HandleFunc("/users/{id:[0-9]+}", userHandler.DeleteUser).Methods("DELETE")
    r.HandleFunc("/login", userHandler.Login).Methods("POST") // Add login route

    // Serve the registration HTML template
    r.HandleFunc("/register-page", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "../../static/register.html")
    })

    // Start server
    log.Println("Starting user service on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", r))
}

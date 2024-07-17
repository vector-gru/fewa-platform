// main.go

package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/lesi/tutor_booking_system/pkg/database"
    "github.com/lesi/tutor_booking_system/handlers"
    "github.com/lesi/tutor_booking_system/services"
)

func main() {
    database.InitDB()

    userService := services.NewUserService(database.DB)
    userHandler := handlers.NewUserHandler(*userService)

    r := mux.NewRouter()

    r.HandleFunc("/register", userHandler.RegisterUser).Methods("POST")
    r.HandleFunc("/users", userHandler.GetAllUsers).Methods("GET")
    r.HandleFunc("/users/{id}", userHandler.GetUserByID).Methods("GET")
    r.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
    r.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

    log.Println("Starting user service on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", r))
}

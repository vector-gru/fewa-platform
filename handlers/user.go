package handlers

import (
    "context"
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/lesi/tutor_booking_system/models"
    "github.com/lesi/tutor_booking_system/services"
    "github.com/lesi/tutor_booking_system/pkg/database"
    "github.com/lesi/tutor_booking_system/validators"
)

func RegisterUserHandlers(r *mux.Router) {
    db := database.DB  // Assuming DB is initialized in database package
    userService := services.NewUserService(db)

    r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
        handleUsers(w, r, userService)
    }).Methods("GET", "POST")
    
    r.HandleFunc("/index.html", serveIndex).Methods("GET")
    
    fs := http.FileServer(http.Dir("../../static"))
    r.PathPrefix("/").Handler(http.StripPrefix("/", fs))
}

func handleUsers(w http.ResponseWriter, r *http.Request, userService *services.UserService) {
    if r.Method == http.MethodGet {
        getUsers(w, r, userService)
    } else if r.Method == http.MethodPost {
        createUser(w, r, userService)
    } else {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func getUsers(w http.ResponseWriter, r *http.Request, userService *services.UserService) {
    users, err := userService.GetAllUsers(context.Background())
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request, userService *services.UserService) {
    var user models.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    err = validators.ValidateUser(user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    hashedPassword, err := services.HashPassword(user.Password)
    if err != nil {
        http.Error(w, "Failed to hash password", http.StatusInternalServerError)
        return
    }
    user.Password = hashedPassword

    err = userService.CreateUser(context.Background(), &user)
    if err != nil {
        http.Error(w, "Failed to create user", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "../../static/index.html")
}

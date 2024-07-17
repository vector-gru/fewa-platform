package handlers

import (
    "encoding/json"
    "log"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "github.com/lesi/tutor_booking_system/models"
    "github.com/lesi/tutor_booking_system/services"
)

type UserHandler struct {
    userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
    return &UserHandler{
        userService: userService,
    }
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
    var user models.User

    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&user); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        log.Printf("Error decoding JSON: %v", err)
        return
    }
    defer r.Body.Close()

    if user.FirstName == "" || user.LastName == "" || user.Email == "" || user.Password == "" {
        http.Error(w, "Missing required fields", http.StatusBadRequest)
        log.Println("Missing required fields")
        return
    }

    hashedPassword, err := services.HashPassword(user.Password)
    if err != nil {
        http.Error(w, "Failed to hash password", http.StatusInternalServerError)
        log.Printf("Error hashing password: %v", err)
        return
    }
    user.Password = hashedPassword

    if err := h.userService.CreateUser(r.Context(), &user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        log.Printf("Error creating user: %v", err)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)

    log.Printf("User registered successfully: %v", user)
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
    users, err := h.userService.GetAllUsers(r.Context())
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    idStr := vars["id"]
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    user, err := h.userService.GetUserByID(r.Context(), uint(id))
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(user)
}

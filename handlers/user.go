// handlers/user.go

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
    userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
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
        log.Printf("Error getting all users: %v", err)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        log.Printf("Invalid user ID: %v", err)
        return
    }

    user, err := h.userService.GetUserByID(r.Context(), uint(id))
    if err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        log.Printf("User not found: %v", err)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        log.Printf("Invalid user ID: %v", err)
        return
    }

    var user models.User
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&user); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        log.Printf("Error decoding JSON: %v", err)
        return
    }
    defer r.Body.Close()

    user.ID = uint(id)
    if err := h.userService.UpdateUser(r.Context(), &user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        log.Printf("Error updating user: %v", err)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(user)

    log.Printf("User updated successfully: %v", user)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        log.Printf("Invalid user ID: %v", err)
        return
    }

    if err := h.userService.DeleteUser(r.Context(), uint(id)); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        log.Printf("Error deleting user: %v", err)
        return
    }

    w.WriteHeader(http.StatusNoContent)
    log.Printf("User deleted successfully: %v", id)
}

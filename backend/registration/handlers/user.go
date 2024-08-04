package handlers

import (
    "encoding/json"
    "log"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "github.com/lesi/tutor_booking_system/registration/models"
    "github.com/lesi/tutor_booking_system/registration/services"
    "github.com/lesi/tutor_booking_system/registration/utils"
)

type UserHandler struct {
    userService services.UserService
    authService services.AuthService
}

func NewUserHandler(userService services.UserService, authService services.AuthService) *UserHandler {
    return &UserHandler{
        userService: userService,
        authService: authService,
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

    // Validate required fields
    if user.FirstName == "" || user.LastName == "" || user.Email == "" || user.Password == "" || user.Role == "" {
        http.Error(w, "Missing required fields", http.StatusBadRequest)
        log.Println("Missing required fields")
        return
    }

    // Hash the password before saving
    hashedPassword, err := utils.HashPassword(user.Password)
    if err != nil {
        http.Error(w, "Failed to hash password", http.StatusInternalServerError)
        log.Printf("Error hashing password: %v", err)
        return
    }
    user.Password = hashedPassword

    // Create user
    if err := h.userService.CreateUser(&user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        log.Printf("Error creating user: %v", err)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
    log.Printf("User registered successfully: %v", user)
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
    users, err := h.userService.GetAllUsers()
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

    user, err := h.userService.GetUserByID(uint(id))
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

    // Optionally hash the password if it is provided in the request
    if user.Password != "" {
        hashedPassword, err := utils.HashPassword(user.Password)
        if err != nil {
            http.Error(w, "Failed to hash password", http.StatusInternalServerError)
            log.Printf("Error hashing password: %v", err)
            return
        }
        user.Password = hashedPassword
    }

    // Update user
    if err := h.userService.UpdateUser(&user); err != nil {
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

    if err := h.userService.DeleteUser(uint(id)); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        log.Printf("Error deleting user: %v", err)
        return
    }

    w.WriteHeader(http.StatusNoContent)
    log.Printf("User deleted successfully: %v", id)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
    var loginRequest struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&loginRequest); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        log.Printf("Error decoding JSON: %v", err)
        return
    }
    defer r.Body.Close()

    user, err := h.authService.AuthenticateUser(loginRequest.Email, loginRequest.Password)
    if err != nil {
        http.Error(w, "Invalid email or password", http.StatusUnauthorized)
        log.Printf("Authentication failed: %v", err)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

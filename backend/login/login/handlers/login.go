package handlers

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/lesi/tutor_booking_system/login/services"
)

type LoginHandler struct {
    authService services.AuthService
}

func NewLoginHandler(authService services.AuthService) *LoginHandler {
    return &LoginHandler{
        authService: authService,
    }
}

type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

func (h *LoginHandler) Login(w http.ResponseWriter, r *http.Request) {
    var loginRequest LoginRequest
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&loginRequest); err != nil {
        http.Error(w, `{"message": "Invalid request payload"}`, http.StatusBadRequest)
        log.Printf("Error decoding JSON: %v", err)
        return
    }
    defer r.Body.Close()

    user, err := h.authService.AuthenticateUser(loginRequest.Email, loginRequest.Password)
    if err != nil {
        http.Error(w, `{"message": "Invalid email or password"}`, http.StatusUnauthorized)
        log.Printf("Authentication failed for email %s: %v", loginRequest.Email, err)
        return
    }

    // Log successful login
    log.Printf("User logged in successfully: %s", loginRequest.Email)

    // Send success response
    w.Header().Set("Content-Type", "application/json")
    response := map[string]interface{}{
        "message": "Login successful",
        "user":    user, // Include user details if necessary
    }
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, `{"message": "Failed to encode response"}`, http.StatusInternalServerError)
        log.Printf("Error encoding JSON response: %v", err)
    }
}

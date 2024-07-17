package handlers

import (
    "encoding/json"
    "log"
    "net/http"
    "strconv"

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

    // Decode JSON request body into user struct
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&user); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        log.Printf("Error decoding JSON: %v", err)
        return
    }
    defer r.Body.Close()

    // Convert agree_to_terms string to boolean
    agreeToTerms := user.AgreeToTerms == "true" // Assumes "true" or "false" string values
    user.AgreeToTerms = strconv.FormatBool(agreeToTerms)

    // Validate required fields
    if user.FirstName == "" || user.LastName == "" || user.Email == "" || user.Password == "" {
        http.Error(w, "Missing required fields", http.StatusBadRequest)
        log.Println("Missing required fields")
        return
    }

    // Optionally, you can hash the user's password before saving it
    hashedPassword, err := services.HashPassword(user.Password)
    if err != nil {
        http.Error(w, "Failed to hash password", http.StatusInternalServerError)
        log.Printf("Error hashing password: %v", err)
        return
    }
    user.Password = hashedPassword

    // Create user in the database
    if err := h.userService.CreateUser(r.Context(), &user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        log.Printf("Error creating user: %v", err)
        return
    }

    // Return success response
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)

    // Log successful registration
    log.Printf("User registered successfully: %v", user)
}




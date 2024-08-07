package handlers

import (
    "encoding/json"
    "log"
    "io"
    "bytes"
    "net/http"
    "strconv"
    "strings"

    "github.com/gorilla/mux"
    "github.com/lesi/tutor_booking_system/registration/models"
    "github.com/lesi/tutor_booking_system/registration/services"
    "github.com/lesi/tutor_booking_system/registration/utils"
)

type UserHandler struct {
    userService services.UserService
    authService services.AuthService
}

// NewUserHandler creates a new UserHandler instance
func NewUserHandler(userService services.UserService, authService services.AuthService) *UserHandler {
    return &UserHandler{
        userService: userService,
        authService: authService,
    }
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
    var user models.User

    log.Println("RegisterUser handler called")

    // Log the raw request body
    rawBody, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Failed to read request body", http.StatusInternalServerError)
        log.Printf("Error reading request body: %v", err)
        return
    }

    // Log the raw request body for debugging
    log.Printf("Raw request body: %s", rawBody)

    // Decode the request body into the user struct
    decoder := json.NewDecoder(bytes.NewReader(rawBody))
    if err := decoder.Decode(&user); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        log.Printf("Error decoding JSON: %v", err)
        return
    }

    // Log received user data for debugging
    log.Printf("Decoded user data: %+v", user)

    // Validate required fields
    missingFields := []string{}
    if user.FirstName == "" {
        missingFields = append(missingFields, "first_name")
    }
    if user.LastName == "" {
        missingFields = append(missingFields, "last_name")
    }
    if user.Email == "" {
        missingFields = append(missingFields, "email")
    }
    if user.Password == "" {
        missingFields = append(missingFields, "password")
    }
    if user.Role == "" {
        missingFields = append(missingFields, "role")
    }

    if len(missingFields) > 0 {
        errorMsg := "Missing required fields: " + strings.Join(missingFields, ", ")
        http.Error(w, errorMsg, http.StatusBadRequest)
        log.Printf("Missing required fields: %v", missingFields)
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

    // Log hashed password for debugging (Remove in production)
    log.Printf("Hashed password: %s", user.Password)

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



// GetAllUsers retrieves all users
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

// GetUserByID retrieves a user by ID
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

// UpdateUser updates a user's details
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

// DeleteUser deletes a user by ID
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

// Login handles user login
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
    log.Println("Login handler called") // Log that the handler is invoked

    var loginRequest struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    
    // Decode the request body
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&loginRequest); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        log.Printf("Error decoding JSON: %v", err)
        return
    }
    defer r.Body.Close()

    log.Printf("Decoded login request: Email=%s", loginRequest.Email) // Log decoded request data

    // Authenticate the user
    user, err := h.authService.AuthenticateUser(loginRequest.Email, loginRequest.Password)
    if err != nil {
        http.Error(w, "Invalid email or password", http.StatusUnauthorized)
        log.Printf("Authentication failed for email %s: %v", loginRequest.Email, err)
        return
    }

    // Log successful login
    log.Printf("User logged in successfully: %s", loginRequest.Email)

    // Send success response
    w.Header().Set("Content-Type", "application/json")
    response := map[string]interface{}{
        "message": "Login successful",
        "user":    user,
    }
    if err := json.NewEncoder(w).Encode(response); err != nil {
        log.Printf("Error encoding response: %v", err)
    }
}

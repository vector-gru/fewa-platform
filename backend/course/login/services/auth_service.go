package services

import (
    "bytes"
    "encoding/json"
    "errors"
    "io"
    "net/http"

    "github.com/lesi/tutor_booking_system/registration/models"
)

// AuthService interface for authentication
type AuthService interface {
    AuthenticateUser(email, password string) (*models.User, error)
}

type authService struct {
    registrationServiceURL string
}

// NewAuthService creates a new instance of authService
func NewAuthService(registrationServiceURL string) AuthService {
    return &authService{
        registrationServiceURL: registrationServiceURL,
    }
}

// AuthenticateUser sends a login request to the registration service
func (s *authService) AuthenticateUser(email, password string) (*models.User, error) {
    loginRequest := map[string]string{
        "email":    email,
        "password": password,
    }
    requestBody, err := json.Marshal(loginRequest)
    if err != nil {
        return nil, err
    }

    resp, err := http.Post(s.registrationServiceURL+"/authenticate", "application/json", bytes.NewBuffer(requestBody))
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        // Read response body to provide more detailed error information
        responseBody, _ := io.ReadAll(resp.Body)
        return nil, errors.New("authentication failed: " + string(responseBody))
    }

    var user models.User
    if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
        return nil, err
    }

    return &user, nil
}

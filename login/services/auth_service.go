package services

import (
    "fmt"
    "net/http"
    "encoding/json"
    "io/ioutil"
)

type AuthService interface {
    AuthenticateUser(email, password string) (map[string]interface{}, error)
}

type authService struct {
    registrationServiceURL string
}

func NewAuthService(registrationServiceURL string) AuthService {
    return &authService{
        registrationServiceURL: registrationServiceURL,
    }
}

// Correct method receiver for the authService struct
func (s *authService) AuthenticateUser(email, password string) (map[string]interface{}, error) {
    // Prepare the request to the registration service
    req, err := http.NewRequest("GET", fmt.Sprintf("%s/users?email=%s", s.registrationServiceURL, email), nil)
    if err != nil {
        return nil, err
    }
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("failed to retrieve user")
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var users []map[string]interface{}
    if err := json.Unmarshal(body, &users); err != nil {
        return nil, err
    }

    if len(users) == 0 {
        return nil, fmt.Errorf("user not found")
    }

    user := users[0]
    storedPassword, ok := user["password"].(string)
    if !ok {
        return nil, fmt.Errorf("password field missing in user data")
    }

    // Compare the provided password with the stored password
    if password != storedPassword {
        return nil, fmt.Errorf("invalid password")
    }

    return user, nil
}

package user

import (
    "context"
    "sync"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

type Service interface {
    GetAllUsers(ctx context.Context) ([]User, error)
}

type userService struct {
    users []User
    mu    sync.RWMutex
}

func NewService() Service {
    return &userService{
        users: []User{
            {ID: 1, Name: "John Doe", Email: "john.doe@example.com"},
            {ID: 2, Name: "Jane Smith", Email: "jane.smith@example.com"},
        },
    }
}

func (s *userService) GetAllUsers(ctx context.Context) ([]User, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    return s.users, nil
}

package course

import (
    "context"
    "sync"
)

type Course struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
}

type Service interface {
    GetAllCourses(ctx context.Context) ([]Course, error)
}

type courseService struct {
    courses []Course
    mu      sync.RWMutex
}

func NewService() Service {
    return &courseService{
        courses: []Course{
            {ID: 1, Name: "Go Programming", Description: "Learn the Go programming language."},
            {ID: 2, Name: "Microservices", Description: "Learn about microservices architecture."},
            {ID: 3, Name: "Kubernetes", Description: "Learn Kubernetes programming and Helm charts."},
        },
    }
}

func (s *courseService) GetAllCourses(ctx context.Context) ([]Course, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    return s.courses, nil
}

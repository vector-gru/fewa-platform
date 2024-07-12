package course

import (
    "context"
    "github.com/lesi/tutor_booking_system/pkg/database"
    "github.com/lesi/tutor_booking_system/models"
)

type Service interface {
    GetAllCourses(ctx context.Context) ([]models.Course, error)
}

type courseService struct{}

func NewService() Service {
    return &courseService{}
}

func (s *courseService) GetAllCourses(ctx context.Context) ([]models.Course, error) {
    var courses []models.Course
    result := database.DB.Find(&courses)
    return courses, result.Error
}

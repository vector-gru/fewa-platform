package user

import (
    "context"
    "github.com/lesi/tutor_booking_system/pkg/database"
    "github.com/lesi/tutor_booking_system/models"
)

type Service interface {
    GetAllUsers(ctx context.Context) ([]models.User, error)
}

type userService struct{}

func NewService() Service {
    return &userService{}
}

func (s *userService) GetAllUsers(ctx context.Context) ([]models.User, error) {
    var users []models.User
    result := database.DB.Find(&users)
    return users, result.Error
}

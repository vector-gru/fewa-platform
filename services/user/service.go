package user

import (
    "context"
    "log"
    "github.com/lesi/tutor_booking_system/pkg/database"
    "github.com/lesi/tutor_booking_system/models"
)

type Service interface {
    GetAllUsers(ctx context.Context) ([]models.User, error)
}

type userService struct {}

func NewService() Service {
    return &userService{}
}

func (s *userService) GetAllUsers(ctx context.Context) ([]models.User, error) {
    var users []models.User
    if err := database.DB.Find(&users).Error; err != nil {
        log.Println("Error fetching users:", err)
        return nil, err
    }
    log.Println("Fetched users:", users)
    return users, nil
}

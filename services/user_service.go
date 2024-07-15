package services

import (
    "context"
    "github.com/lesi/tutor_booking_system/models"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
)

type UserService struct {
    db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
    return &UserService{db: db}
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]models.User, error) {
    var users []models.User
    err := s.db.WithContext(ctx).Find(&users).Error
    return users, err
}

func (s *UserService) CreateUser(ctx context.Context, user *models.User) error {
    return s.db.WithContext(ctx).Create(user).Error
}

func HashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}

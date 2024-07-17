package services

import (
    "context"

    "github.com/lesi/tutor_booking_system/models"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
)

// UserService interface defines the methods that a user service should implement.
type UserService interface {
    GetAllUsers(ctx context.Context) ([]models.User, error)
    CreateUser(ctx context.Context, user *models.User) error
}

type userService struct {
    db *gorm.DB
}

// NewUserService returns a new instance of UserService.
func NewUserService(db *gorm.DB) UserService {
    return &userService{db: db}
}

func (s *userService) GetAllUsers(ctx context.Context) ([]models.User, error) {
    var users []models.User
    err := s.db.WithContext(ctx).Find(&users).Error
    return users, err
}

func (s *userService) CreateUser(ctx context.Context, user *models.User) error {
    hashedPassword, err := HashPassword(user.Password)
    if err != nil {
        return err
    }
    user.Password = hashedPassword
    return s.db.WithContext(ctx).Create(user).Error
}

func HashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}

// services/user_service.go

package services

import (
    "context"
    "github.com/lesi/tutor_booking_system/models"
    "gorm.io/gorm"
    "golang.org/x/crypto/bcrypt"
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

func (s *UserService) GetUserByID(ctx context.Context, id uint) (*models.User, error) {
    var user models.User
    err := s.db.WithContext(ctx).First(&user, id).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (s *UserService) CreateUser(ctx context.Context, user *models.User) error {
    return s.db.WithContext(ctx).Create(user).Error
}

func (s *UserService) UpdateUser(ctx context.Context, user *models.User) error {
    return s.db.WithContext(ctx).Save(user).Error
}

func (s *UserService) DeleteUser(ctx context.Context, id uint) error {
    return s.db.WithContext(ctx).Delete(&models.User{}, id).Error
}

func HashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}

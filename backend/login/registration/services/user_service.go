package services

import (
    "github.com/lesi/tutor_booking_system/registration/models"
    "gorm.io/gorm"
    "golang.org/x/crypto/bcrypt"
)

// UserService interface for managing users
type UserService interface {
    CreateUser(user *models.User) error
    GetAllUsers() ([]models.User, error)
    GetUserByID(id uint) (*models.User, error)
    UpdateUser(user *models.User) error
    DeleteUser(id uint) error
}

// AuthService interface for authentication
type AuthService interface {
    AuthenticateUser(email, password string) (*models.User, error)
}

// userService struct implementing UserService
type userService struct {
    db *gorm.DB
}

// authService struct implementing AuthService
type authService struct {
    db *gorm.DB
}

// NewUserService creates a new instance of userService
func NewUserService(db *gorm.DB) UserService {
    return &userService{db: db}
}

// NewAuthService creates a new instance of authService
func NewAuthService(db *gorm.DB) AuthService {
    return &authService{db: db}
}

// UserService methods

func (s *userService) CreateUser(user *models.User) error {
    // Hash the password before saving the user
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(hashedPassword)

    return s.db.Create(user).Error
}

func (s *userService) GetAllUsers() ([]models.User, error) {
    var users []models.User
    err := s.db.Find(&users).Error
    return users, err
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
    var user models.User
    err := s.db.First(&user, id).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (s *userService) UpdateUser(user *models.User) error {
    return s.db.Save(user).Error
}

func (s *userService) DeleteUser(id uint) error {
    return s.db.Delete(&models.User{}, id).Error
}

// AuthService methods

func (s *authService) AuthenticateUser(email, password string) (*models.User, error) {
    var user models.User
    err := s.db.Where("email = ?", email).First(&user).Error
    if err != nil {
        return nil, err
    }

    // Compare hashed password
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return nil, err
    }

    return &user, nil
}

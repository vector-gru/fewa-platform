package course

import (
    "context"
    "github.com/lesi/tutor_booking_system/pkg/logging"
    "github.com/lesi/tutor_booking_system/models"
    "gorm.io/gorm"
)

type Service interface {
    GetAllCourses(ctx context.Context) ([]models.Course, error)
}

type service struct {
    db     *gorm.DB
    logger logging.Logger
}

func NewService(db *gorm.DB, logger logging.Logger) Service {
    return &service{
        db:     db,
        logger: logger,
    }
}

func (s *service) GetAllCourses(ctx context.Context) ([]models.Course, error) {
    var courses []models.Course
    result := s.db.Find(&courses)
    if result.Error != nil {
        s.logger.Error("Failed to fetch courses:", result.Error)
        return nil, result.Error
    }
    return courses, nil
}

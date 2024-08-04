package repositories

import (
    "course_service/models"
    "gorm.io/gorm"
)

type CourseRepository struct {
    DB *gorm.DB
}

func NewCourseRepository(db *gorm.DB) *CourseRepository {
    return &CourseRepository{DB: db}
}

func (r *CourseRepository) CreateCourse(course *models.Course) error {
    return r.DB.Create(course).Error
}

func (r *CourseRepository) GetAllCourses() ([]models.Course, error) {
    var courses []models.Course
    err := r.DB.Find(&courses).Error
    return courses, err
}

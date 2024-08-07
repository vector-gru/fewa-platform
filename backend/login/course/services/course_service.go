package services

import (
    "course_service/models"
    "course_service/repositories"
)

type CourseService struct {
    Repo *repositories.CourseRepository
}

func NewCourseService(repo *repositories.CourseRepository) *CourseService {
    return &CourseService{Repo: repo}
}

func (s *CourseService) CreateCourse(course *models.Course) error {
    return s.Repo.CreateCourse(course)
}

func (s *CourseService) GetAllCourses() ([]models.Course, error) {
    return s.Repo.GetAllCourses()
}

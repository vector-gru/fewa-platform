package controllers

import (
    "encoding/json"
    "net/http"
    "course_service/models"
    "course_service/services"
)

type CourseController struct {
    Service *services.CourseService
}

func NewCourseController(service *services.CourseService) *CourseController {
    return &CourseController{Service: service}
}

func (c *CourseController) CreateCourse(w http.ResponseWriter, r *http.Request) {
    var course models.Course
    json.NewDecoder(r.Body).Decode(&course)
    err := c.Service.CreateCourse(&course)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(course)
}

func (c *CourseController) GetAllCourses(w http.ResponseWriter, r *http.Request) {
    courses, err := c.Service.GetAllCourses()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(courses)
}

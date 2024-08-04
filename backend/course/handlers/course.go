package handlers

import (
    "course_service/models"
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    "gorm.io/gorm"
)

func RegisterCourseRoutes(router *mux.Router, db *gorm.DB) {
    router.HandleFunc("/courses", CreateCourse(db)).Methods("POST")
    router.HandleFunc("/courses/{id}", GetCourse(db)).Methods("GET")
    router.HandleFunc("/courses", GetAllCourses(db)).Methods("GET")
    router.HandleFunc("/courses/{id}", UpdateCourse(db)).Methods("PUT")
    router.HandleFunc("/courses/{id}", DeleteCourse(db)).Methods("DELETE")
}

func CreateCourse(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var course models.Course
        if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        if err := db.Create(&course).Error; err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(course)
    }
}

func GetCourse(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        params := mux.Vars(r)
        var course models.Course
        if err := db.First(&course, params["id"]).Error; err != nil {
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }
        json.NewEncoder(w).Encode(course)
    }
}

func GetAllCourses(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var courses []models.Course
        if err := db.Find(&courses).Error; err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        json.NewEncoder(w).Encode(courses)
    }
}

func UpdateCourse(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        params := mux.Vars(r)
        var course models.Course
        if err := db.First(&course, params["id"]).Error; err != nil {
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }
        var updatedCourse models.Course
        if err := json.NewDecoder(r.Body).Decode(&updatedCourse); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        db.Model(&course).Updates(updatedCourse)
        json.NewEncoder(w).Encode(course)
    }
}

func DeleteCourse(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        params := mux.Vars(r)
        var course models.Course
        if err := db.First(&course, params["id"]).Error; err != nil {
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }
        db.Delete(&course)
        w.WriteHeader(http.StatusNoContent)
    }
}

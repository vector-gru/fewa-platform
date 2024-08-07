package routes

import (
    "course_service/controllers"
    "github.com/gorilla/mux"
)

func RegisterCourseRoutes(router *mux.Router, controller *controllers.CourseController) {
    router.HandleFunc("/courses", controller.CreateCourse).Methods("POST")
    router.HandleFunc("/courses", controller.GetAllCourses).Methods("GET")
}

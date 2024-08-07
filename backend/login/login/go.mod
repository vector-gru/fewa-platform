module github.com/lesi/tutor_booking_system/login

go 1.22.5

require (
	github.com/gorilla/mux v1.8.0
	github.com/lesi/tutor_booking_system/registration v0.0.0
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/crypto v0.25.0 // indirect
	gorm.io/gorm v1.21.9 // indirect
)

replace github.com/lesi/tutor_booking_system/registration => ../registration

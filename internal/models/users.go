package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	first_name string
	last_name  string
	email      string
	password   string
	phone      string
}

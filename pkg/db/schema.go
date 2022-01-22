package db

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	TeacherID uint   `json:"teacher_id" gorm:"not null;primaryKey;unique;"`
	FullName  string `json:"full_name" gorm:"size:255;not null;"`
	Email     string `json:"email" gorm:"size:120;"`
}

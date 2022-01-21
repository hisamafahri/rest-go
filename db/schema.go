package db

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model

	TeacherID uint   `gorm:"PrimaryKey;unique_index"`
	FullName  string `gorm:"typevarchar(255);not null"`
	Email     string `gorm:"typevarchar(100)"`
}

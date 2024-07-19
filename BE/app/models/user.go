package models

type User struct {
	Email    string `gorm:"size:255;not null;uniqueIndex;primary_key"`
	Password string `gorm:"size:255;not null"`
}

package models

import "time"

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Name     string `gorm:"size:255"`
	Email    string `gorm:"size:255;unique"`
	Pwd      string `gorm:"size:255"`
	Birthday *time.Time 
}

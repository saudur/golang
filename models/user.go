package models

import (
	time "time"
)

type User struct {
	Id        int8      `gorm:"primaryKey; autoIncrement:true" json:"id"`
	Name      string    `gorm:"size:255;not null;" json:"name"`
	Password  string    `gorm:"size:255;not null;" json:"password"`
	CreatedAt time.Time `gorm:"type:timestamp;not null;default:now()" json:"created"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null;default:now()" json:"update"`
}

package models

import "time"

type Author struct {
	Id        uint   `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"type:varchar(255)" json:"name"`
	Gender    string `gorm:"type:char(1)" json:"gender"`
	Email     string `gorm:"type:varchar(255)" json:"email"`
	Age       int    `gorm:"type:integer" json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

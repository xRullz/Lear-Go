package entity

import "time"

type User struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255)" json:"name"`
	Address   string    `gorm:"type:text" json:"address"`
	Email     string    `gorm:"type:varchar(255)" json:"email"`
	Phone     string    `gorm:"type:varchar(255)" json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

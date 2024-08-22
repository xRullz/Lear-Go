package entity

import "time"

type User struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255)" json:"name"`
	Address   string    `gorm:"type:text" json:"address"`
	Phone     string    `gorm:"type:varchar(255)" json:"phone"`
	Email     string    `gorm:"type:varchar(255)" json:"email"`
	Password  string    `gorm:"column:password" json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

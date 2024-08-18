package request

type UserCreateRequest struct {
	Name      string    `gorm:"type:varchar(255)" json:"name"`
	Address   string    `gorm:"type:text" json:"address"`
	Email     string    `gorm:"type:varchar(255)" json:"email"`
	Phone     string    `gorm:"type:varchar(255)" json:"phone"`
}
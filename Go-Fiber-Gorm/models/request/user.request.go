package request

type UserCreateRequest struct {
	Name      string    `gorm:"type:varchar(255)" json:"name" validate:"required"`
	Address   string    `gorm:"type:text" json:"address"`
	Email     string    `gorm:"type:varchar(255)" json:"email" validate:"required,email"`
	Phone     string    `gorm:"type:varchar(255)" json:"phone"`
	Password  string    `gorm:"type:varchar(255)" json:"password" validate:"required,min=5"`

}

type UserUpdateRequest struct {
	Name      string    `gorm:"type:varchar(255)" json:"name"`
	Address   string    `gorm:"type:text" json:"address"`
	Phone     string    `gorm:"type:varchar(255)" json:"phone"`
	Email     string    `gorm:"type:varchar(255)" json:"email"`
	Password  string    `gorm:"type:varchar(255)" json:"password"`

}
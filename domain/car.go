package domain

type Car struct {
	ID                 int    `gorm:"primary_key"`
	Brand              string `json:"brand"`
	Color              string `json:"color"`
	Model              string `json:"model"`
	ModelYear          int    `json:"model_year"`
	Price              int    `json:"price"`
	RegistrationNumber string `json:"registration_number"`

	UserID int  `json:"user_id"`
	User   User `json:"user"`
}

package domain

type User struct {
	ID   int    `gorm:"primary_key"`
	Name string `json:"name"`
	Cars []Car  `json:"cars" gorm:"foreignKey:UserID"`
}

package model

type User struct {
	ID          int    `json:"id"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	Password    string `json:"password"`
	CategoryID  int    `json:"category_id" db:"category_id"`
}

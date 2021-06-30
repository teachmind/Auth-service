package model

type User struct {
	ID          int    `json:"id"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	Password    string `json:"password"`
	CategoryId  int    `json:"category_id"`
}

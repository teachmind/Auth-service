package model

type User struct {
<<<<<<< HEAD
	ID           int    `json:"id"`
	Phone        string `json:"phone"`
	FullName     string `json:"full_name" db:"full_name"`
	Password     string `json:"password"`
	BusinessName string `json:"business_name" db:"business_name"`
=======
	ID          int    `json:"id"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	Password    string `json:"password"`
	CategoryId  int    `json:"category_id" db:"category_id"`
>>>>>>> main
}

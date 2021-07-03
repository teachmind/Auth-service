package model

import (
	"regexp"
)

type User struct {
	ID          int    `json:"id"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	Password    string `json:"password"`
	CategoryId  int    `json:"category_id" db:"category_id"`
}

func PhoneValidation(phone_number string) bool {
	re := regexp.MustCompile(`^(\+88)?(01)(\d{3})[ -]?(\d{6})$`)
	err := re.MatchString(phone_number)
	return err
}

package model

import (
	"fmt"
	"regexp"
)

type User struct {
	ID          int    `json:"id"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	Password    string `json:"password"`
	CategoryID  int    `json:"category_id" db:"category_id"`
}

// Regex patter
const (
	regexpPhone = `^(\+88)?(01)(\d{3})[ -]?(\d{6})$`
)

// ValidateLogin Validates user login input
func (u User) ValidateLogin() error {

	if u.PhoneNumber == "" {
		return fmt.Errorf("Phone Number is required :%w", ErrEmpty)
	}

	if len(u.PhoneNumber) != 14 {
		return fmt.Errorf("Phone Number must be 14 digit :%w", ErrEmpty)
	}

	if !u.PhoneValidation() {
		return fmt.Errorf("Phone Number must start with +880. Ex: +8801712345678 :%w", ErrInvalid)
	}

	if u.Password == "" {
		return fmt.Errorf("Password is required :%w", ErrEmpty)
	}

	return nil
}

// PhoneValidation Validate phone numbers with regexp
func (u User) PhoneValidation() bool {
	re := regexp.MustCompile(regexpPhone)
	err := re.MatchString(u.PhoneNumber)
	return err
}

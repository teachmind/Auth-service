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

// ValidAuthentication Validates user input credentials
func (u *User) ValidateAuthentication() error {

	if u.PhoneNumber == "" {
		return fmt.Errorf("Phone Number is required :%w", ErrEmpty)
	}

	if !u.PhoneValidation() {
		return fmt.Errorf("Phone Number should look like this 01712345678 :%w", ErrInvalid)
	}

	if u.Password == "" {
		return fmt.Errorf("Password is required :%w", ErrEmpty)
	}

	return nil
}

// PhoneValidation Validate phone numbers with regexp
func (u User) PhoneValidation() bool {
	re := regexp.MustCompile(regexpPhone)
	isValid := re.MatchString(u.PhoneNumber)
	return isValid
}

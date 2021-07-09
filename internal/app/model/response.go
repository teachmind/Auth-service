package model

import (
	"errors"
	"fmt"
)

var ErrNotFound = fmt.Errorf("not found")
var ErrInvalid = fmt.Errorf("invalid")
var ErrEmpty = fmt.Errorf("empty") // Empty error
var ErrUnauthorized = errors.New("access denied")

type GenericResponse struct {
	Success bool                   `json:"success"`
	Errors  []ErrorDetailsResponse `json:"errors"`
	Data    interface{}            `json:"data"`
}

type ErrorDetailsResponse struct {
	Code     string `json:"code,omitempty"`
	Message  string `json:"message,omitempty"`
	Title    string `json:"message_title,omitempty"`
	Severity string `json:"severity,omitempty"`
}

// LoginResponse returns response after successful login
type LoginResponse struct {
	ID          int    `json:"id"`
	PhoneNumber string `json:"phone_number"`
	CategoryID  int    `json:"category_id"`
	Token       string `json:"token"`
}

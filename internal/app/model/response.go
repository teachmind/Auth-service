package model

import (
	"errors"
	"fmt"
)

var ErrNotFound = fmt.Errorf("not found")
var ErrInvalid = fmt.Errorf("invalid")
var ErrEmpty = fmt.Errorf("empty")
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

type LoginResponse struct {
	ID          int    `json:"id"`
	PhoneNumber string `json:"phone_number"`
	CategoryId  int    `json:"category_id"`
	Token       string `json:"token"`
}

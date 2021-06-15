package user

import (
	"github.com/lib/pq"
)

const (
	errUniqueViolation = pq.ErrorCode("23505")
)

type repository struct {
}

func NewRepository() *repository {
	return &repository{}
}

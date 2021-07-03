package user

import (
	"context"
	"database/sql"
	"fmt"
	"user-service/internal/app/model"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

// SQL Query and error
const (
	errUniqueViolation        = pq.ErrorCode("23505")
	GetUserByPhoneNumberQuery = `SELECT id, phone_number, password FROM users WHERE phone_number = $1`
)

type repository struct {
	db *sqlx.DB
}

// NewRepository initiates user repository and returns DB
func NewRepository(db *sqlx.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetUserByPhoneNumber(ctx context.Context, PhoneNumber string) (model.User, error) {
	var user model.User
	if err := r.db.GetContext(ctx, &user, GetUserByPhoneNumberQuery, PhoneNumber); err != nil {
		if err == sql.ErrNoRows {
			return model.User{}, fmt.Errorf("No record found for this phone number. :%w", model.ErrNotFound)
		}
		return model.User{}, err
	}
	return user, nil
}

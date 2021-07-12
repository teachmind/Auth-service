package user

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"user-service/internal/app/model"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

// SQL Query and error
const (
	errUniqueViolation  = pq.ErrorCode("23505")
	insertUserQuery     = `INSERT INTO users (phone_number, password, category_id) VALUES ($1, $2, $3)`
	getUserByPhoneQuery = `SELECT id, phone_number, password, category_id FROM users WHERE phone_number = $1`
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

func (r *repository) InsertUser(ctx context.Context, user model.User) error {
	phoneNumber := RMCodeAndSpace(user.PhoneNumber)
	if _, err := r.db.ExecContext(ctx, insertUserQuery, phoneNumber, user.Password, 1); err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == errUniqueViolation {
			return fmt.Errorf("%v :%w", err, model.ErrInvalid)
		}
		return err
	}
	return nil
}

func (r *repository) GetUserByPhone(ctx context.Context, phoneNumber string) (model.User, error) {
	var user model.User
	actualPhone := RMCodeAndSpace(phoneNumber)
	if err := r.db.GetContext(ctx, &user, getUserByPhoneQuery, actualPhone); err != nil {
		if err == sql.ErrNoRows {
			return model.User{}, fmt.Errorf("the phone no is not found :%w", model.ErrNotFound)
		}
		return model.User{}, err
	}
	return user, nil
}

// RMCodeAndSpace remove the country code and space from http request phone no
func RMCodeAndSpace (phoneNumber string) (string) {
	phoneNumber = strings.ReplaceAll(phoneNumber, " ", "")
	phoneNumber = strings.ReplaceAll(phoneNumber, "+88", "")
	return phoneNumber
}

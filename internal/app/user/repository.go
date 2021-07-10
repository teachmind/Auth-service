package user

import (
	"context"
	"database/sql"
	"fmt"
	"user-service/internal/app/model"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

const (
	errUniqueViolation  = pq.ErrorCode("23505")
	insertUserQuery     = `INSERT INTO users (phone_number, password, category_id) VALUES ($1, $2, $3)`
	getUserByPhoneQuery = `SELECT id, phone_number, password, category_id FROM users WHERE phone_number = $1`
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) InsertUser(ctx context.Context, user model.User) error {
	if _, err := r.db.ExecContext(ctx, insertUserQuery, user.PhoneNumber, user.Password, 1); err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == errUniqueViolation {
			return fmt.Errorf("%v :%w", err, model.ErrInvalid)
		}
		return err
	}
	return nil
}

func (r *repository) GetUserByPhone(ctx context.Context, phoneNumber string) (model.User, error) {
	var user model.User
	if err := r.db.GetContext(ctx, &user, getUserByPhoneQuery, phoneNumber); err != nil {
		if err == sql.ErrNoRows {
			return model.User{}, fmt.Errorf("the phone no is not found :%w", model.ErrNotFound)
		}
		return model.User{}, err
	}
	return user, nil
}

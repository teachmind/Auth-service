package user

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"user-service/internal/app/model"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestRepository_GetUserByPhoneNumber(t *testing.T) {
	t.Run("should return success", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		user := model.User{
			ID:          1,
			PhoneNumber: "123456",
			CategoryId:  1,
			Password:    "123456",
		}
		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectQuery("^SELECT (.+) FROM users WHERE (.+)").
			WithArgs("phone_number").
			WillReturnRows(sqlmock.NewRows([]string{"id", "phone_number", "category_id", "password"}).
				AddRow(1, "123456", 1, "123456"))
		repo := NewRepository(sqlxDB)
		result, err := repo.GetUserByPhoneNumber(context.Background(), "phone_number")
		assert.Nil(t, err)
		assert.EqualValues(t, user, result)
	})

	t.Run("should return no rows error", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectQuery("^SELECT (.+) FROM users WHERE (.+)").
			WithArgs("phone_number").
			WillReturnError(sql.ErrNoRows)
		repo := NewRepository(sqlxDB)
		_, err := repo.GetUserByPhoneNumber(context.Background(), "phone_number")
		assert.True(t, errors.Is(err, model.ErrNotFound))
	})

	t.Run("should return error", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectQuery("^SELECT (.+) FROM users WHERE (.+)").
			WithArgs("phone_number").
			WillReturnError(errors.New("sql-error"))
		repo := NewRepository(sqlxDB)
		_, err := repo.GetUserByPhoneNumber(context.Background(), "phone_number")
		assert.NotNil(t, err)
	})
}

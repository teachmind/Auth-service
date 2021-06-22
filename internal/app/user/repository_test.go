package user

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"user-service/internal/app/model"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestRepository_InsertUser(t *testing.T) {
	t.Run("should return success", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		user := model.User{
			Phone:        "phone",
			FullName:     "full_name",
			Password:     "password",
			BusinessName: "business_name",
		}
		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectExec("INSERT INTO users (.+) VALUES (.+)").
			WithArgs(user.Phone, user.FullName, user.Password, user.BusinessName).
			WillReturnResult(sqlmock.NewResult(1, 1))

		repo := NewRepository(sqlxDB)
		err := repo.InsertUser(context.Background(), user)
		assert.True(t, errors.Is(err, model.ErrInvalid))
	})

	t.Run("should return unique key violation error", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		user := model.User{
			Phone:        "phone",
			FullName:     "full_name",
			Password:     "password",
			BusinessName: "business_name",
		}

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectExec("INSERT INTO users (.+) VALUES (.+)").
			WithArgs(user.Phone, user.FullName, user.Password, user.BusinessName).
			WillReturnError(&pq.Error{Code: "23505"})

		repo := NewRepository(sqlxDB)
		err := repo.InsertUser(context.Background(), user)
		assert.True(t, errors.Is(err, model.ErrInvalid))
	})

	t.Run("should return sql error", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		user := model.User{
			Phone:        "phone",
			FullName:     "full_name",
			Password:     "password",
			BusinessName: "business_name",
		}

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectExec("INSERT INTO users (.+) VALUES (.+)").
			WithArgs(user.Phone, user.FullName, user.Password, user.BusinessName).
			WillReturnError(errors.New("sql-error"))

		repo := NewRepository(sqlxDB)
		err := repo.InsertUser(context.Background(), user)
		assert.NotNil(t, err)
	})
}

func TestRepository_GetUserByPhone(t *testing.T) {
	t.Run("should return success", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		user := model.User{
			ID:           1,
			Phone:        "01738799349",
			FullName:     "Mr. Name",
			Password:     "123456",
			BusinessName: "business-1",
		}

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectQuery("^SELECT (.+) FROM users WHERE (.+)").
			WithArgs("phone").
			WillReturnRows(sqlmock.NewRows([]string{"id", "phone", "full_name", "password", "business_name"}).
				AddRow(1, "01738799349", "Mr. Name", "123456", "business-1"))

		repo := NewRepository(sqlxDB)
		result, err := repo.GetUserByPhone(context.Background(), "phone")
		assert.Nil(t, err)
		assert.EqualValues(t, user, result)
	})

	t.Run("should return no rows error", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectQuery("^SELECT (.+) FROM users WHERE (.+)").
			WithArgs("phone").
			WillReturnError(sql.ErrNoRows)
		repo := NewRepository(sqlxDB)
		_, err := repo.GetUserByPhone(context.Background(), "phone")
		assert.True(t, errors.Is(err, model.ErrNotFound))
	})

	t.Run("should return error", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectQuery("^SELECT (.+) FROM users WHERE (.+)").
			WithArgs("phone").
			WillReturnError(errors.New("sql-error"))
		repo := NewRepository(sqlxDB)
		_, err := repo.GetUserByPhone(context.Background(), "phone")
		assert.NotNil(t, err)
	})
}

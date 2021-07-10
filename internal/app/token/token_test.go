package token

import (
	"testing"
	"user-service/internal/app/model"

	"github.com/stretchr/testify/assert"
)

func TestService_Decode(t *testing.T) {
	t.Run("should decode the same user", func(t *testing.T) {
		user := model.User{
			ID:          2,
			PhoneNumber: "+8801707123123",
			Password:    "pas8889ff",
			CategoryId:  1,
		}

		s := NewService()
		token, err := s.Encode(user)
		assert.Nil(t, err)
		jwtClm, err := s.Decode(token)
		assert.Nil(t, err)
		assert.EqualValues(t, user, jwtClm.User)
	})

	t.Run("should return token parse error", func(t *testing.T) {
		token := ""
		s := NewService()
		jwtClm, err := s.Decode(token)
		assert.NotNil(t, err)
		assert.Nil(t, jwtClm)
	})
}

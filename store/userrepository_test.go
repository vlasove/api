package store_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vlasove/api/internal/app/model"
	"github.com/vlasove/api/store"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, config)
	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email: "test_user@example.org",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, config)
	defer teardown("users")

	email := "test_user@example.org"
	t.Run("not existing user", func(t *testing.T) {
		u, err := s.User().FindByEmail(email)
		assert.Error(t, err)
		assert.Nil(t, u)
	})

	t.Run("existing user", func(t *testing.T) {
		_, err := s.User().Create(&model.User{
			Email: email,
		})
		assert.NoError(t, err)

		u, err := s.User().FindByEmail(email)
		assert.NoError(t, err)
		assert.NotNil(t, u)

	})
}

package teststore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vlasove/api/internal/app/model"
	"github.com/vlasove/api/internal/app/store"
	"github.com/vlasove/api/internal/app/store/teststore"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()

	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.New()

	testUser := model.TestUser(t)
	email := "test_user@example.org"
	testUser.Email = email
	t.Run("not existing user", func(t *testing.T) {
		u, err := s.User().FindByEmail(email)
		assert.EqualError(t, err, store.ErrRecordNotFound.Error())
		assert.Nil(t, u)
	})

	t.Run("existing user", func(t *testing.T) {
		err := s.User().Create(testUser)
		assert.NoError(t, err)

		u, err := s.User().FindByEmail(email)
		assert.NoError(t, err)
		assert.NotNil(t, u)

	})
}

func TestUserRepository_Find(t *testing.T) {
	s := teststore.New()
	testUser := model.TestUser(t)
	err := s.User().Create(testUser)
	assert.NoError(t, err)
	u, err := s.User().Find(testUser.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vlasove/api/internal/app/model"
)

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "empty email",
			u: func() *model.User {
				user := model.TestUser(t)
				user.Email = ""
				return user
			},
			isValid: false,
		},
		{
			name: "not an email",
			u: func() *model.User {
				user := model.TestUser(t)
				user.Email = "ivalidemail"
				return user
			},
			isValid: false,
		},
		{
			name: "empty password",
			u: func() *model.User {
				user := model.TestUser(t)
				user.Password = ""
				return user
			},
			isValid: false,
		},
		{
			name: "short password",
			u: func() *model.User {
				user := model.TestUser(t)
				user.Password = "abc"
				return user
			},
			isValid: false,
		},
		{
			name: "with encrypted password",
			u: func() *model.User {
				user := model.TestUser(t)
				user.Password = ""
				user.EncryptedPassword = "encryptedpassword"
				return user
			},
			isValid: true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isValid {
				assert.NoError(t, tt.u().Validate())
			} else {
				assert.Error(t, tt.u().Validate())
			}
		})
	}
}

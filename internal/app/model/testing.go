package model

import "testing"

// TestUser ...
func TestUser(t *testing.T) *User {
	t.Helper()
	return &User{
		Email:    "user_test@example.org",
		Password: "password",
	}

}

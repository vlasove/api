// Package model shoud describe data model in project
package model

// User ...
type User struct {
	ID                int
	Email             string
	EncryptedPassword string
}

package models

import (
	"time"
)

type User struct {
	Username       string
	Name           string
	Email          string
	PictureSlug    string
	Description    string
	HashedPassword []byte
}

type Session struct {
	Token  []byte
	Expiry time.Time
	User   *User
}

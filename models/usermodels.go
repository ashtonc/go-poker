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
	HashedPassword string
}

type Session struct {
	Token  string
	Expiry time.Time
	User   *User
}

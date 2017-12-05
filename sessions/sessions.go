package sessions

import (
	"poker/models"
)

// GetSession returns a session object with information about the current user.
func GetSession(id string) *models.Session {
	var session models.Session

	session.LoggedIn = true
	session.Username = "current-user"
	session.Name = "Current User"

	return &session
}
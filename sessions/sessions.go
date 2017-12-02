package sessions

import (
	"poker/models"
)

// GetSession returns a session object with information about the current user. It should probably be passed a session ID that would be provided by the user.
func GetSession() models.Session {
	var session models.Session

	session.LoggedIn = true
	session.Username = "current-user"
	session.Name = "Current User"

	return session
}

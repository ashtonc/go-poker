package sessions

import (
	// "time"

	"poker/database"
	"poker/models"
)

// GetSession returns a session object with information about the current user.
func GetSessionWithInfo(env *models.Env, id string) *models.Session {
	session, err := database.GetSession(env, id)
	if err != nil {
		// session not found,
		session.LoggedIn = false
	}

	// if session.Expiry .. is past time.Now().
	//			session.LoggedIn = false

	// Debugging values until sessions actually exist in the database
	session.LoggedIn = true
	session.Username = "current-user"

	return session
}
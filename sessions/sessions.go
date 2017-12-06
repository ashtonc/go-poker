package sessions

import (
	_ "time"
	"fmt"

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

	// Debugging values until sessions actually exist in the database...
	session.LoggedIn = true
	session.Username = "current-user"

	return session
}

// CreateSession creates a new session with a random token and saves it in the database.
// It returns the token string and an error.
// func CreateSession(env *models.Env, username string, expiry time.Time) (string, error) {
func CreateSession(env *models.Env, username string) *models.Session {

	var session models.Session
	//session = sessions.GetSessionWithInfo(env, "blah")
	userAccount := database.FindByUsername(env, username)

	session.UserId = (database.GetUserID(env, username)).UserId
	session.Username = username
	session.Name = userAccount.Name
	session.Email = userAccount.Email
	session.LoggedIn = true

	fmt.Printf("Session values: \n", session.Id, session.Username, session.Name, session.Email, session.UserId)
	err := database.NewSession(env, "token", '18:00', session.UserId)
	if err != nil {
		// session not found,
		session.LoggedIn = false
	}

	return &session
}

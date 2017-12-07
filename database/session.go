package database

import (
	"time"

	"github.com/gorilla/securecookie"

	"poker/models"
)

func GetSession(env *models.Env, sessionid string) (*models.Session, error) {
	var session models.Session
	var user models.User
	var userid int

	sqlStatement := `SELECT token, expiry_time, user_id FROM user_session WHERE token = $1`
	row := env.Database.QueryRow(sqlStatement, sessionid)
	err := row.Scan(&session.Token, &session.Expiry, &userid)
	if err != nil {
		return nil, err
	}

	// Check the expiry time here, return an error if expired

	sqlStatement = `SELECT username, name, email, picture_slug, description FROM account WHERE id = $1`
	row = env.Database.QueryRow(sqlStatement, userid)
	err = row.Scan(&user.Username, &user.Name, &user.Email, &user.PictureSlug, &user.Description)
	if err != nil {
		return nil, err
	}

	session.User = &user
	return &session, err
}

func CreateSession(env *models.Env, username string, expiry time.Time) (string, error) {
	var userid int
	token := securecookie.GenerateRandomKey(512)

	sqlStatement := `SELECT id FROM account WHERE username = $1`
	row := env.Database.QueryRow(sqlStatement, username)
	err := row.Scan(&userid)
	if err != nil {
		return token, err
	}

	sqlStatement = `INSERT INTO user_session (token, expiry_time, user_id) VALUES ($1, $2, $3)`
	_, err = env.Database.Exec(sqlStatement, token, expiry, userid)
	if err != nil {
		return token, err
	}

	return token, err
}

func RevokeSession(env *models.Env, username string) error {
	var userid int

	sqlStatement := `SELECT id FROM account WHERE username = $1`
	row := env.Database.QueryRow(sqlStatement, username)
	err := row.Scan(&userid)
	if err != nil {
		return err
	}

	sqlStatement = `DELETE FROM user_session WHERE user_id = $1`
	_, err = env.Database.Exec(sqlStatement, userid)
	if err != nil {
		return err
	}

	return err
}

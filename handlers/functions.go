package handlers

import (
	"log"
	"net/http"

	"poker/database"
	"poker/models"
)

func getPageData(env *models.Env, r *http.Request, sessionid string, page string) models.PageData {
	var pagedata models.PageData
	pagedata.SiteRoot = env.SiteRoot

	switch page {
	case "Home":
		pagedata.NavigationLevel = models.NAVIGATION_HOME
	case "Game", "Lobby":
		pagedata.NavigationLevel = models.NAVIGATION_GAME
	case "Leaderboard":
		pagedata.NavigationLevel = models.NAVIGATION_LEADERBOARD
	case "ViewUser", "EditUser", "Login", "Register":
		pagedata.NavigationLevel = models.NAVIGATION_USER
	case "Admin":
		pagedata.NavigationLevel = models.NAVIGATION_ADMIN
	}

	var identity models.Identity
	pagedata.Identity = &identity
	pagedata.Identity.LoggedIn = false

	token, err := getSessionToken(env, r)
	if err != nil {
		log.Print(err)
		return pagedata
	}

	session, err := database.GetSession(env, token)
	if err != nil {
		log.Print(err)
		return pagedata
	}

	pagedata.Identity.LoggedIn = true
	pagedata.Identity.AccountType = models.TYPE_USER_ACCOUNT
	pagedata.Identity.Username = session.User.Username
	pagedata.Identity.Name = session.User.Name

	return pagedata
}

func setSessionToken(env *models.Env, w http.ResponseWriter, r *http.Request, token string) error {
	value := map[string]string{
		"token": token,
	}

	encoded, err := env.CookieHandler.Encode("poker-470-session", value)
	if err == nil {
		cookie := &http.Cookie{
			Name:  "poker-470-session",
			Value: encoded,
			Path:  "/",
		}

		http.SetCookie(w, cookie)

		return err
	}

	return err
}

func getSessionToken(env *models.Env, r *http.Request) (string, error) {
	cookie, err := r.Cookie("poker-470-session")
	if err == nil {
		value := make(map[string]string)
		if err = env.CookieHandler.Decode("poker-470-session", cookie.Value, &value); err == nil {
			return value["token"], err
		}

		return "", err
	}

	return "", err
}

func clearSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "poker-470-session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}

	http.SetCookie(w, cookie)
}

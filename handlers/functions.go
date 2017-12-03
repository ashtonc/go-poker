package handlers

import (
	"net/http"

	"github.com/gorilla/securecookie"

	"poker/models"
	"poker/sessions"
)

// Move the global variables in the env struct ************
var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func getPageData(env *models.Env, sessionid string, page string) models.PageData {
	var pagedata models.PageData
	session := sessions.GetSession(sessionid)

	switch page {
	case "Home":
		session.PageHome = true
	case "ViewUser", "EditUser", "Login", "Register":
		session.PageUser = true
	case "Login":
		session.PageLogin = true
	case "Register":
		session.PageRegister = true
	case "PlayGame", "WatchGame", "ViewLobby":
		session.PageGame = true
	case "Leaderboard":
		session.PageLeaderboard = true
	}

	pagedata.Session = session
	pagedata.SiteRoot = env.SiteRoot

	return pagedata
}

func getUserName(request *http.Request) (userName string) {
	if cookie, err := request.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["username"]
		}
	}
	return userName
}

func getName(request *http.Request) (name string) {
	if cookie, err := request.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			name = cookieValue["name"]
		}
	}
	return name
}

func setSession(userName string, name string, response http.ResponseWriter) {
	value := map[string]string{
		"username": userName,
		"name":     name,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

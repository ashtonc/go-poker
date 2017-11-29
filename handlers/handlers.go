package handlers

import (
	"log"
	"net/http"
	"html/template"
	"fmt"

	"poker/models"
	"poker/database"

	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
)

// This simply redirects users to /poker
func HomeRedirect(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/poker/", http.StatusTemporaryRedirect)
	})
}

func Home(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Populate the data needed for the page (these should nearly all be external functions)
		pagedata := models.PageData{
			Session: models.Session{
				LoggedIn: true,
				Username: "current-user",
				Name: "Current User",
				PageHome: true,
			},
		}

		// Build our template using the required files (need base, head, navigation, and content)
		// This should be moved to a caching function: https://elithrar.github.io/article/approximating-html-template-inheritance/
		t, _ := template.ParseFiles("./templates/base.tmpl", "./templates/head_base.tmpl", "./templates/navigation.tmpl", "./templates/index.tmpl")

		// Execute the template with our page data
		t.Execute(w, pagedata)
	})
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


func clearSession(response http.ResponseWriter) {
    cookie := &http.Cookie{
        Name:   "session",
        Value:  "",
        Path:   "/",
        MaxAge: -1,
    }
    http.SetCookie(response, cookie)
}


var cookieHandler = securecookie.New(
    securecookie.GenerateRandomKey(64),
    securecookie.GenerateRandomKey(32))

func setSession(userName string, name string, response http.ResponseWriter) {
    value := map[string]string{
        "username": userName,
        "name": name,
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

func Login(env *models.Env) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
// func Login(response http.ResponseWriter, request *http.Request) {


		if request.Method == "POST" {
			fmt.Printf("TTTTT")
			userName := request.FormValue("username")
			name := request.FormValue("username")
	     pass := request.FormValue("password")
	     redirectTarget := "/poker/login/"
	     if name != "" && pass != "" {

			fmt.Printf("NAME")
	        // .. check credentials ..
	        setSession(userName, name, response)
	        redirectTarget = "/poker/game/"
	     }
	     //redirect to "404 page not found"
	     http.Redirect(response, request, redirectTarget, 302)
		} else {
	// Populate the data needed for the page (these should nearly all be external functions)
			fmt.Printf("after if")
			pagedata := models.PageData{
				Session: models.Session{
					LoggedIn: false,
					PageLogin: true,
				},
			}
			// Build our template using the required files (need base, head, navigation, and content)
			// This should be moved to a caching function: https://elithrar.github.io/article/approximating-html-template-inheritance/
			t, _ := template.ParseFiles("./templates/base.tmpl", "./templates/head_base.tmpl", "./templates/navigation.tmpl", "./templates/login.tmpl")

			// Execute the template with our page data
			t.Execute(response, pagedata)
		}

		})
	
}


func logoutHandler(response http.ResponseWriter, request *http.Request) {
    clearSession(response)
    http.Redirect(response, request, "/", 302)
}

func Register(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET" {
			// Populate the data needed for the page (these should nearly all be external functions)
			pagedata := models.PageData{
				Session: models.Session{
					LoggedIn: false,
					PageUser: true,
				},
			}

			// Build our template using the required files (need base, head, navigation, and content)
			// This should be moved to a caching function: https://elithrar.github.io/article/approximating-html-template-inheritance/
			t, _ := template.ParseFiles("./templates/base.tmpl", "./templates/head_base.tmpl", "./templates/navigation.tmpl", "./templates/register.tmpl")

			// Execute the template with our page data
			t.Execute(w, pagedata)
		} else if r.Method == "POST" {
			fmt.Printf("test");

		}
	})
}

func User(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		username := vars["username"]

		// Get the user page matching that username from the database
		user, err := database.GetUserPage(env, username)
		if err != nil {
			log.Print("Player " + username + " not found.")

			// For now, just redirect them to the home page
			http.Redirect(w, r, "/poker/", http.StatusTemporaryRedirect)
			return
		}

		// Populate the data needed for the page (these should nearly all be external functions)
		pagedata := models.PageData{
			Session: models.Session{
				LoggedIn: true,
				Username: "current-user",
				Name: "Current User",
				PageUser: true,
			},
			UserPage: models.UserPage{
				MatchesSession: true,
				Username: user.Username,
				Name: user.Name,
				Email: user.Email,
				PictureUrl: user.PictureUrl,
			},
		}

		// Build our template using the required files (need base, head, navigation, and content)
		// This should be moved to a caching function: https://elithrar.github.io/article/approximating-html-template-inheritance/
		t, _ := template.ParseFiles("./templates/base.tmpl", "./templates/head_base.tmpl", "./templates/navigation.tmpl", "./templates/user_view.tmpl")

		log.Print("Displaying player " + username + ".")
		// Execute the template with our page data
		t.Execute(w, pagedata)
	})
}

func UserEdit(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		username := vars["username"]

		// Populate the data needed for the page (these should nearly all be external functions)
		pagedata := models.PageData{
			Session: models.Session{
				LoggedIn: true,
				Username: "current-user",
				Name: "Current User",
				PageUser: true,
			},
			UserPage: models.UserPage{
				MatchesSession: true,
				Username: username,
				Name: "User Name",
				Email: "user@email.ca",
			},
		}

		// Build our template using the required files (need base, head, navigation, and content)
		// This should be moved to a caching function: https://elithrar.github.io/article/approximating-html-template-inheritance/
		t, _ := template.ParseFiles("./templates/base.tmpl", "./templates/head_base.tmpl", "./templates/navigation.tmpl", "./templates/user_edit.tmpl")

		// Execute the template with our page data
		t.Execute(w, pagedata)
	})
}

func Game(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Populate the data needed for the page (these should nearly all be external functions)
		pagedata := models.PageData{
			Session: models.Session{
				LoggedIn: true,
				Username: getUserName(r),
				Name: getName(r),
				PageGame: true,
			},
		}

		// Build our template using the required files (need base, head, navigation, and content)
		// This should be moved to a caching function: https://elithrar.github.io/article/approximating-html-template-inheritance/
		t, _ := template.ParseFiles("./templates/base.tmpl", "./templates/head_base.tmpl", "./templates/navigation.tmpl", "./templates/game_lobby.tmpl")

		// Execute the template with our page data
		t.Execute(w, pagedata)
	})
}

func Leaderboard(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Populate the data needed for the page (these should nearly all be external functions)
		pagedata := models.PageData{
			Session: models.Session{
				LoggedIn: true,
				Username: "current-user",
				Name: "Current User",
				PageLeaderboard: true,
			},
		}

		// Build our template using the required files (need base, head, navigation, and content)
		// This should be moved to a caching function: https://elithrar.github.io/article/approximating-html-template-inheritance/
		t, _ := template.ParseFiles("./templates/base.tmpl", "./templates/head_base.tmpl", "./templates/navigation.tmpl", "./templates/leaderboard.tmpl")

		// Execute the template with our page data
		t.Execute(w, pagedata)
	})
}

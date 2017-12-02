package templates

import (
	"html/template"
	"log"
)

func BuildTemplates() map[string]*template.Template {

	home, err := template.ParseFiles("./templates/base.tmpl", "./templates/head_base.tmpl", "./templates/navigation.tmpl", "./templates/index.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	login, err := template.ParseFiles("./templates/base.tmpl", "./templates/head_base.tmpl", "./templates/navigation.tmpl", "./templates/login.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	register, err := template.ParseFiles("./templates/base.tmpl", "./templates/head_base.tmpl", "./templates/navigation.tmpl", "./templates/register.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	viewUser, err := template.ParseFiles("./templates/base.tmpl", "./templates/head_base.tmpl", "./templates/navigation.tmpl", "./templates/user_view.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	editUser, err := template.ParseFiles("./templates/base.tmpl", "./templates/head_base.tmpl", "./templates/navigation.tmpl", "./templates/user_edit.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	playGame, err := template.ParseFiles("./templates/base.tmpl", "./templates/head_game.tmpl", "./templates/navigation.tmpl", "./templates/game_play.tmpl", "./templates/game.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	viewLobby, err := template.ParseFiles("./templates/base.tmpl", "./templates/head_base.tmpl", "./templates/navigation.tmpl", "./templates/game_lobby.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	watchGame, err := template.ParseFiles("./templates/base.tmpl", "./templates/head_game.tmpl", "./templates/navigation.tmpl", "./templates/game_watch.tmpl", "./templates/game.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	leaderboard, err := template.ParseFiles("./templates/base.tmpl", "./templates/head_base.tmpl", "./templates/navigation.tmpl", "./templates/leaderboard.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	templateCache := map[string]*template.Template{
		"Home":        home,
		"Login":       login,
		"Register":    register,
		"ViewUser":    viewUser,
		"EditUser":    editUser,
		"PlayGame":    playGame,
		"ViewLobby":   viewLobby,
		"WatchGame":   watchGame,
		"Leaderboard": leaderboard,
	}

	return templateCache
}

package templates

import (
	"html/template"
	"log"
)

func BuildTemplateCache() map[string]*template.Template {
	log.Print("Building the template cache...")

	templateCache := map[string]*template.Template{
		"Home":        BuildWithError("templates/base.tmpl", "templates/head_base.tmpl", "templates/navigation.tmpl", "templates/index.tmpl"),
		"Login":       BuildWithError("templates/base.tmpl", "templates/head_base.tmpl", "templates/navigation.tmpl", "templates/login.tmpl"),
		"Register":    BuildWithError("templates/base.tmpl", "templates/head_base.tmpl", "templates/navigation.tmpl", "templates/register.tmpl"),
		"ViewUser":    BuildWithError("templates/base.tmpl", "templates/head_base.tmpl", "templates/navigation.tmpl", "templates/user_view.tmpl"),
		"EditUser":    BuildWithError("templates/base.tmpl", "templates/head_base.tmpl", "templates/navigation.tmpl", "templates/user_edit.tmpl"),
		"PlayGame":    BuildWithError("templates/base.tmpl", "templates/head_game.tmpl", "templates/navigation.tmpl", "templates/game_play.tmpl", "templates/game.tmpl"),
		"ViewLobby":   BuildWithError("templates/base.tmpl", "templates/head_base.tmpl", "templates/navigation.tmpl", "templates/game_lobby.tmpl"),
		"WatchGame":   BuildWithError("templates/base.tmpl", "templates/head_game.tmpl", "templates/navigation.tmpl", "templates/game_watch.tmpl", "templates/game.tmpl"),
		"Leaderboard": BuildWithError("templates/base.tmpl", "templates/head_base.tmpl", "templates/navigation.tmpl", "templates/leaderboard.tmpl"),
	}

	return templateCache
}

// Similar to the functionality of template.Must, but writes the error to the server log
func BuildWithError(filenames ...string) *template.Template {
	template, err := template.ParseFiles(filenames...)
	if err != nil {
		log.Fatal(err)
	}

	return template
}

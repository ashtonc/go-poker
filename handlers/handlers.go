package handlers

import (
	"poker-project/models"
	"net/http"
)

func HomeRedirect(env *models.Env) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/poker", http.StatusTemporaryRedirect)
	})
}

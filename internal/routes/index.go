package routes

import (
	"net/http"
	"sendit/internal/context"
)

func HandleIndex(w http.ResponseWriter, r *http.Request, app *context.App) {
	data := struct {
		DisplayName string
		DisplayDescription string
	}{
		DisplayName: app.Config.DisplayName,
		DisplayDescription: app.Config.DisplayDescription,
	}
	err := app.IndexTpl.Execute(w, data)
	if err != nil {
		panic(err)
	}
}


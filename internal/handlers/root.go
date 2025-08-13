package handlers

import (
	"github.com/Meschkov/htmx-playground/web"
	"net/http"

	"html/template"
	"log/slog"
)

type PageData struct {
	Title string
}

func RootHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file, err := web.Templates.ReadFile("templates/index.html")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			slog.Error("Error reading template", "error", err)
			return
		}

		tmpl := template.Must(template.New("index.html").Parse(string(file)))

		data := PageData{
			Title: "Home",
		}
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			slog.Error("Error executing template", "error", err)
		}
	}
}

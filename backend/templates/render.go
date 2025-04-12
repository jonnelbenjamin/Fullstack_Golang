package templates

import (
	"html/template"
	"net/http"
	"time"
)

var templates *template.Template

func InitTemplates() error {
	var err error
	templates, err = template.New("").Funcs(template.FuncMap{
		"formatDate": func(t time.Time) string {
			return t.Format("Jan 02, 2006")
		},
	}).ParseGlob("templates/*.html")
	return err
}

func Render(w http.ResponseWriter, name string, data interface{}) {
	err := templates.ExecuteTemplate(w, name, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

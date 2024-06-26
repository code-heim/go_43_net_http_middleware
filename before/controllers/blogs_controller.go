package controllers

import (
	"net/http"
	"net_http_middleware/models"
	"path"
	"text/template"
)

func BlogsIndex(w http.ResponseWriter, r *http.Request) {
	blogs := models.BlogsAll()
	fp := path.Join("templates", "blogs", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, blogs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

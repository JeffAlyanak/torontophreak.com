package ui

import (
	"net/http"
	"html/template"
	"toronto-phreak/config"
	"toronto-phreak/model"
)

var Templates *template.Template

func CacheTemplates() error {
	var err error
	Templates, err	= template.ParseGlob(config.Conf.AssetDir + "templates/index.html")
	return err
}

func RenderTemplate(w http.ResponseWriter, tstr string, al model.ArticleList) {
	err := Templates.ExecuteTemplate(w, "layout", al)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
}

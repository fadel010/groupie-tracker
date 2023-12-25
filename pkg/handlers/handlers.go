package handlers

import (
	"groupie-tracker/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplateHome(w, r, "home.html")
}

func Info(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplateInfo(w, r, "info.html")
}

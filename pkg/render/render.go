package render

import (
	"fmt"
	"groupie-tracker/pkg/api"
	"groupie-tracker/pkg/model"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

func RenderTemplateHome(w http.ResponseWriter, r *http.Request, tmpl string) {
	if r.URL.Path != "/" {
		RenderErrorPage(w, http.StatusNotFound)
		return
	}
	artists, err := api.GetArtists()
	if err != nil {
		log.Printf("Error in GetArtists: %v", err)
		RenderErrorPage(w, http.StatusInternalServerError)
		return
	}

	parsedTemplate, _ := loadTemplate("templates/" + tmpl)
	err = parsedTemplate.Execute(w, artists)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		RenderErrorPage(w, http.StatusInternalServerError)
		return
	}
}

func RenderTemplateInfo(w http.ResponseWriter, r *http.Request, tmpl string) {
	path := r.URL.Path
	a := strings.Split(path, "/info/")
	i, err := strconv.Atoi(strings.Join(a, ""))

	if err != nil {
		RenderErrorPage(w, http.StatusNotFound)
		return
	}

	if i <= 0 || i >= 53 {
		log.Printf("Invalid artist ID: %v", err)
		RenderErrorPage(w, http.StatusBadRequest)
		return
	}

	artist, err := api.GetArtistByID(i)
	if err != nil {
		log.Printf("Error in GetArtistByID: %v", err)
		RenderErrorPage(w, http.StatusInternalServerError)
		return
	}

	concertDate, err := api.GetDateByID(i)
	if err != nil {
		log.Printf("Error in GetDateByID: %v", err)
		RenderErrorPage(w, http.StatusInternalServerError)
		return
	}
	for i := range concertDate.Dates {
		concertDate.Dates[i] = strings.ReplaceAll(concertDate.Dates[i],"*", "")
	}

	locations, err := api.GetLocationByID(i)
	if err != nil {
		log.Printf("Error in GetLocationByID: %v", err)
		RenderErrorPage(w, http.StatusInternalServerError)
		return
	}
	relations, err := api.GetRelationByID(i)
	if err != nil {
		log.Printf("Error in GetRelationByID: %v", err)
		RenderErrorPage(w, http.StatusInternalServerError)
		return
	}

	allInfo := model.AllInfo{
		Artist:         artist,
		Locations:      locations,
		ConcertDates:   concertDate,
		Relations:      relations,
	}

	// log.Println("AllInfo Data:", allInfo)

	parsedTemplate, _ := loadTemplate("templates/" + tmpl)
	err = parsedTemplate.Execute(w, allInfo)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		RenderErrorPage(w, http.StatusInternalServerError)
		return
	}
}

func loadTemplate(filename string) (*template.Template, error) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		return nil, fmt.Errorf("error parsing template %s: %v", filename, err)
	}
	return tmpl, nil
}

func handleError(w http.ResponseWriter, statusCode int) {
	http.Error(w, fmt.Sprintf("%d %s", statusCode, http.StatusText(statusCode)), statusCode)
}

func RenderErrorPage(w http.ResponseWriter, statusCode int) {
	tmpl, err := loadTemplate(fmt.Sprintf("templates/%d.html", statusCode))
	if err != nil {
		log.Printf("Error loading template: %v", err)
		handleError(w, statusCode)
		return
	}
	w.WriteHeader(statusCode)
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		handleError(w, statusCode)
		return
	}
}

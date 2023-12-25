package api

import (
	"encoding/json"
	"groupie-tracker/pkg/model"
	"net/http"
	"strconv"
)

func GetArtists() ([]model.Artist, error) {
	url := ApiLinks.Artists
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var artists []model.Artist
	err = json.NewDecoder(response.Body).Decode(&artists)
	if err != nil {
		return nil, err
	}

	return artists, nil
}

func GetArtistByID(artistID int) (model.Artist, error) {
	url := ApiLinks.Artists + "/" + strconv.Itoa(artistID)
	response, err := http.Get(url)
	if err != nil {
		return model.Artist{}, err
	}
	defer response.Body.Close()
	var artist model.Artist
	err = json.NewDecoder(response.Body).Decode(&artist)
	if err != nil {
		return model.Artist{}, err
	}

	return artist, nil
}

func GetDateByID(artistID int) (model.Date, error) {
	url := ApiLinks.Dates + "/" + strconv.Itoa(artistID)
	response, err := http.Get(url)
	if err != nil {
		return model.Date{}, err
	}
	defer response.Body.Close()
	var date model.Date
	err = json.NewDecoder(response.Body).Decode(&date)
	if err != nil {
		return model.Date{}, err
	}

	return date, nil
}

func GetLocationByID(artistID int) (model.Location, error) {
	url := ApiLinks.Locations + "/" + strconv.Itoa(artistID)
	response, err := http.Get(url)
	if err != nil {
		return model.Location{}, err
	}
	defer response.Body.Close()
	var location model.Location
	err = json.NewDecoder(response.Body).Decode(&location)
	if err != nil {
		return model.Location{}, err
	}

	return location, nil
}

func GetRelationByID(artistID int) (model.Relation, error) {
	url := ApiLinks.Relation + "/" + strconv.Itoa(artistID)
	response, err := http.Get(url)
	if err != nil {
		return model.Relation{}, err
	}
	defer response.Body.Close()
	var relation model.Relation
	err = json.NewDecoder(response.Body).Decode(&relation)
	if err != nil {
		return model.Relation{}, err
	}

	return relation, nil
}

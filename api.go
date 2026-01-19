package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

const BaseURL = "https://groupietrackers.herokuapp.com/api"

func fetchData(endpoint string, target interface{}) error {
	resp, err := http.Get(BaseURL + endpoint)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}

func GetArtists() ([]Artist, error) {
	var artists []Artist
	err := fetchData("/artists", &artists)
	return artists, err
}

func GetLocations(id int) (Locations, error) {
	var locs Locations
	err := fetchData("/locations/"+strconv.Itoa(id), &locs)
	return locs, err
}

func GetDates(id int) (Dates, error) {
	var dates Dates
	err := fetchData("/dates/"+strconv.Itoa(id), &dates)
	return dates, err
}

func GetRelation(id int) (Relation, error) {
	var rel Relation
	err := fetchData("/relation/"+strconv.Itoa(id), &rel)
	return rel, err
}

package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	artists, err := GetArtists()
	if err != nil {
		http.Error(w, "Erreur API", 500)
		return
	}
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, artists)
}

func DetailsHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	artist, _ := fetchDataArtist(id)
	locs, _ := GetLocations(id)
	dates, _ := GetDates(id)
	rel, _ := GetRelation(id)

	data := struct {
		Artist    Artist
		Locations Locations
		Dates     Dates
		Relation  Relation
	}{artist, locs, dates, rel}

	tmpl := template.Must(template.ParseFiles("templates/details.html"))
	tmpl.Execute(w, data)
}

func fetchDataArtist(id int) (Artist, error) {
	var a Artist
	err := fetchData("/artists/"+strconv.Itoa(id), &a)
	return a, err
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/artist", DetailsHandler)

	fmt.Println("Serveur : http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
//new: log.Fatal, html templates
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Film struct {
	Title string
	Director string
	ReleaseYear uint16
}

type FilmList struct {
	ListTitle string
	Films []Film
}

func getFilms(writer http.ResponseWriter, request *http.Request) {
	pageTemplate := template.Must(template.ParseFiles("index.html"))

	bestFilms := FilmList{
		ListTitle: "Best Films Ever",
		Films: []Film{
			{Title: "Mean Girls", Director: "Mark Waters", ReleaseYear: 2004},
			{Title: "Rope", Director: "Alfred Hitchcock", ReleaseYear: 1948},
			{Title: "Night of the Living Dead", Director: "George A. Romero", ReleaseYear: 1968},
		},
	}

	pageTemplate.Execute(writer, bestFilms)
}

func addFilm(writer http.ResponseWriter, request *http.Request) {
	newFilmTitle := request.PostFormValue("title")
	newFilmDirector := request.PostFormValue("director")
	newFilmReleaseYear, _ := strconv.Atoi(request.PostFormValue("releaseyear"))

	htmlString := fmt.Sprintf("<li class='list-group-item bg-primary text-white'>%s - %s, Released %d</li>", newFilmTitle, newFilmDirector, newFilmReleaseYear)
	
	template, _ := template.New("t").Parse(htmlString) //wtf is 't'?

	template.Execute(writer, nil)
}

func main() {
	http.HandleFunc("/", getFilms)
	http.HandleFunc("/add-film/", addFilm)
	log.Fatal(http.ListenAndServe(":8000", nil)) //log.Fatal will write text if there is an issue with the function inside it

}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

//struct is basicly a object in a javascript code
//defining model for incoming movie json to parse 
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"Isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
//defining model for incoming movies director which is associated with movie json to parse 
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
// we are doing it without a database hence we are using a Go slice for the DB work
var movies []Movie

func main() {

	//appending a couple of movies
	movies = append(movies, Movie{ID: "1", Isbn: "1234", Title: "inception", Director: &Director{Firstname: "christopher", Lastname: "Nolan"}})
	movies = append(movies, Movie{ID: "2", Isbn: "12345", Title: "darkKnight", Director: &Director{Firstname: "christopher", Lastname: "Nolan"}})
	//intilaizing the mux router
	r := mux.NewRouter()
	//setting up the handler functions for the incoming requests
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovieById).Methods("GET")
	r.HandleFunc("/movies", createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovies).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovies).Methods("DELETE")

	fmt.Printf("server started at port 8000\n")
	//setting up the server for the localhost
	log.Fatal(http.ListenAndServe(":8000", r))

}

//getMovies is a func which is used to get the list of movies 
func getMovies(w http.ResponseWriter,_*http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

//getMoviesById is a func which is used to get the movieList which as the specific Id that the user given 
func getMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parms := mux.Vars(r)

	for _, item := range movies {
		if item.ID == parms["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}
//createMovies is handler func which is used to create a new movies to the moviesList in the movies slice(DB)
func createMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}
//update the movie which as the specific Id that is given in the request parms
func updateMovies(w http.ResponseWriter, r *http.Request) {
	//set json content type
	w.Header().Set("Content-Type", "application/json")
	//get parms from router
	parms := mux.Vars(r)
	//iterate over the movies
	for index, item := range movies {
		if item.ID == parms["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = parms["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}
// deleteMovies is reqHandler func which is used delete the movie which as same Id that the user given as parms 
func deleteMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parms := mux.Vars(r)

	for index, item := range movies {
		if item.ID == parms["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}






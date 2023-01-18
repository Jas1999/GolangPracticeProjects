package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http" //create server
	"strconv"  // int -> str

	"github.com/gorilla/mux"
)

// struct movie
type Movie struct {
	ID       string    `json: "id"`
	Isbn     string    `json: "isbn"`
	Title    string    `json: "title"`
	Director *Director `json: "director"`
}

// strcut director
type Director struct {
	firstName string `json: "firstName"`
	lastName  string `json: "lastName"`
}

var movies []Movie

// r is pointer to request that will b sent
// w is response item that is sent
func getMovies(w http.ResponseWriter, r *http.Request) {
	// set content type to json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	// set content type to json
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)           //id from front
	for idx, item := range movies { // essentailly for each

		if item.ID == params["id"] {
			movies = append(movies[:idx], movies[idx+1:]...) // delete by skipping idx, append previous and after
			break
		}
	}
	json.NewEncoder(w).Encode(movies) // return remaining movies

}

func getMovie(w http.ResponseWriter, r *http.Request) {
	// set content type to json
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //id from front

	for _, item := range movies { // essentailly for each, need blank since not using index

		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item) // return item
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	// set content type to json
	w.Header().Set("Content-Type", "application/json")
	//params := mux.Vars(r) //id from front

	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)    // read arguments into struct
	movie.ID = strconv.Itoa(rand.Intn(100000000)) // ran between 0 and this value
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie) // return item
	return
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	// set content type to json
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //id from front

	// find movie and delete
	for idx, item := range movies { // essentailly for each

		if item.ID == params["id"] {
			movies = append(movies[:idx], movies[idx+1:]...) // delete by skipping idx, append previous and after
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie) // read arguments into struct
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie) // return item
			return
		}
	}

}

//movies crud w api using gorilla mux
func main() {
	r := mux.NewRouter()
	//dummy data
	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{firstName: "JOHN", lastName: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "436556", Title: "Movie Two", Director: &Director{firstName: "JOHN", lastName: "Doe"}})

	// 5 func for routes
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}

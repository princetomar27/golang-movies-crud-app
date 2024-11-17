package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func GetMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	movieId := strings.TrimPrefix(r.URL.Path,"/movie/")

 

	for _, currentMovie := range Movies{
		if currentMovie.ID == movieId{
			json.NewEncoder(w).Encode(currentMovie)
			return
		}
	}

	http.Error(w,"Movie not found", http.StatusNotFound)
}

func CreateMovie(w http.ResponseWriter, r *http.Request){
	
	w.Header().Set("Content-Type","application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	// Generate movie id
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	// Add movie to Movies Array
	Movies = append(Movies, movie)
	json.NewEncoder(w).Encode(movie)
	 
}


func UpdateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)

	for index, item := range Movies{
		if item.ID == params["id"]{ 
			Movies = append(Movies[:index], Movies[index+1:]... )
			var updatedMovie Movie
			_ =json.NewDecoder(r.Body).Decode(&updatedMovie)
			updatedMovie.ID = params["id"]

			Movies = append(Movies,updatedMovie)
			json.NewEncoder(w).Encode(updatedMovie)
		}
	}
}

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicatin/json")
	json.NewEncoder(w).Encode(Movies)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	
	// Take movie id as a param 
	params := mux.Vars(r)

	for index, item := range Movies{

		if item.ID == params["id"] {
			Movies = append(Movies[:index], Movies[index+1:]... )
			break 
		}
	}

	json.NewEncoder(w).Encode(Movies)
}

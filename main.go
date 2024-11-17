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

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director 	`json:"director"`
	Studio *Studio `json:"studio"`
}

type Director struct{
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

type Studio struct{
	StudioName string `json:"studioname"`
}

var Movies []Movie




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

func GetMovie(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)

	for _, item := range Movies{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func CreateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	var movie Movie

	_ = json.NewDecoder(r.Body).Decode(&movie)

	movie.ID =  strconv.Itoa(rand.Intn(10000000))
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
			json.NewEncoder(w).Encode(Movies)
		}
	}
}

func main(){

	router := mux.NewRouter()

	Movies = append(Movies, Movie{
		ID: "1",
		Isbn: "00112",
		Title: "Kal ho na ho",
		Director: &Director{
			FirstName: "Prince",
			LastName: "Tomar",
		},
		Studio: &Studio{
			StudioName: "Tomar Studios",
		},
	})
	Movies = append(Movies, Movie{
		ID:    "2",
		Isbn:  "00113",
		Title: "Dil Chahta Hai",
		Director: &Director{
			FirstName: "Farhan",
			LastName:  "Akhtar",
		},
		Studio: &Studio{
			StudioName: "Excel Entertainment",
		},
	})	
	Movies = append(Movies, Movie{
		ID:    "3",
		Isbn:  "00114",
		Title: "Lagaan",
		Director: &Director{
			FirstName: "Ashutosh",
			LastName:  "Gowariker",
		},
		Studio: &Studio{
			StudioName: "Aamir Khan Productions",
		},
	})	
	Movies = append(Movies, Movie{
		ID:    "4",
		Isbn:  "00115",
		Title: "3 Idiots",
		Director: &Director{
			FirstName: "Rajkumar",
			LastName:  "Hirani",
		},
		Studio: &Studio{
			StudioName: "Vidhu Vinod Chopra Films",
		},
	})	
	Movies = append(Movies, Movie{
		ID:    "5",
		Isbn:  "00116",
		Title: "Swades",
		Director: &Director{
			FirstName: "Ashutosh",
			LastName:  "Gowariker",
		},
		Studio: &Studio{
			StudioName: "Ashutosh Gowariker Productions",
		},
	})	
	Movies = append(Movies, Movie{
		ID:    "6",
		Isbn:  "00117",
		Title: "Chak De! India",
		Director: &Director{
			FirstName: "Shimit",
			LastName:  "Amin",
		},
		Studio: &Studio{
			StudioName: "Yash Raj Films",
		},
	})
	

	// get all movies
	router.HandleFunc("/movies", GetMovies).Methods("GET")

	// get movie with ID
	router.HandleFunc("/movie/{id}", GetMovie2).Methods("GET")

	// create a movie
	router.HandleFunc("/movies", CreateMovie).Methods("POST")

	// update movie via ID
	router.HandleFunc("/movies/{id}", UpdateMovie).Methods("PUT")

	// delete a movie
	router.HandleFunc("/movie/{id}", DeleteMovie).Methods("DELETE")

	fmt.Printf("Server starting at PORT : 8000")
	log.Fatal(http.ListenAndServe(":8000",router))

}
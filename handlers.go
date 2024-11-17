package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func GetMovie2(w http.ResponseWriter, r *http.Request){
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


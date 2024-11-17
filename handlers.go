package main

import (
	"encoding/json"
	"net/http"
)

func GetMovie2(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	// movieId := strings.TrimPrefix(r.URL.Path,"/movie/")

	movieIdFromQP := r.URL.Query().Get("id")

	if movieIdFromQP == ""{
		http.Error(w, "Movie ID missing in QP", http.StatusBadRequest)
		return
	}

	// params := mux.Vars(r)
	// movieId := params["id"]

	for _, currentMovie := range Movies{
		if currentMovie.ID == movieIdFromQP{
			json.NewEncoder(w).Encode(currentMovie)
			return
		}
	}

	http.Error(w,"Movie not found", http.StatusNotFound)
}


package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)
type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies [] Movie


func getMovies (w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)

}


func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)
	for index ,item :=range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break;
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r*http.Request){
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)
	for _,item := range movies{
		if item.ID==params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie( w http.ResponseWriter, r*http.Request){
	w.Header().Set("Content-Type","application/json")
	movies=append(movies, Movie{ID: "4",Title: "Movie 4", Director: &Director{Firstname: "Sahob",Lastname: "Sab"}})
	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r*http.Request){
	w.Header().Set("Content-Type","application.josn")
	for i,movie:= range movies{
		if movie.Title=="Super Man"{
			movies[i].Title="Avengers"
		}
	}
	json.NewEncoder(w).Encode(movies)
}


func main(){
	r:=mux.NewRouter()

	movies = append(movies,Movie{ID:"1",Title:"Iron Man", Director : &Director{Firstname:"John", Lastname:"Doe"}})
    movies = append(movies,Movie{ID:"2",Title:"Spider Man", Director : &Director{Firstname:"Mir", Lastname:"Pill"}})
    movies = append(movies,Movie{ID:"3",Title:"Super Man", Director : &Director{Firstname:"Mir", Lastname:"Pill"}})
	
	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))
}
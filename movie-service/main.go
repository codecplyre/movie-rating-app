package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
	// "github.com/golang-migrate/migrate/v4"
	// _ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	// _ "github.com/golang-migrate/migrate/v4/source/file"
	// "github.com/gorilla/mux"
)

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(db:3306)/movie")
	if err != nil {
		log.Fatalf("failed to open database connection: %v", err)
	}
	fmt.Println("connected to database")
	defer db.Close()
	fmt.Println("migration successful")

	repo := NewMovieRepository(db)
	service := NewMovieService(repo)

	r := mux.NewRouter()
	r.HandleFunc("/movies", getAllMovies(service)).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovieByID(service)).Methods("GET")
	r.HandleFunc("/movies", createMovie(service)).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie(service)).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie(service)).Methods("DELETE")
	r.HandleFunc("/movies/generate", GenerateFakeData(service)).Methods("POST")

	log.Println("Server is running on port https://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

// fake movie data generator for testing and insert into database
func GenerateFakeData(service MovieService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var movie Movie
		err := json.NewDecoder(r.Body).Decode(&movie)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		for i := 0; i < 100000; i++ {
			movie := Movie{
				Title:       RandomMovie(),
				ReleaseDate: RandomReleaseDate(),
				Director:    RandomDirector(),
			}
			err = service.CreateMovie(&movie)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		json.NewEncoder(w).Encode(movie)
	}
}

func getAllMovies(service MovieService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		movies, err := service.GetMovies()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(movies)
	}
}

func getMovieByID(service MovieService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		movie, err := service.GetMovieByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(movie)
	}
}

func createMovie(service MovieService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var movie Movie
		err := json.NewDecoder(r.Body).Decode(&movie)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = service.CreateMovie(&movie)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(movie)
	}
}

func updateMovie(service MovieService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var movie Movie
		err = json.NewDecoder(r.Body).Decode(&movie)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		movie.ID = id
		err = service.UpdateMovie(&movie)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(movie)
	}
}

func deleteMovie(service MovieService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = service.DeleteMovie(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

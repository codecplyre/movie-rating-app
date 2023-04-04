package main

import "database/sql"

type MovieRepository interface {
	FindAll() ([]*Movie, error)
	FindByID(id int) (*Movie, error)
	Create(movie *Movie) error
	Update(movie *Movie) error
	Delete(id int) error
}

type movieRepository struct {
	db *sql.DB
}

func NewMovieRepository(db *sql.DB) MovieRepository {
	return &movieRepository{db}
}

func (r *movieRepository) FindAll() ([]*Movie, error) {
	// Query the database and return a list of movies
	row, err := r.db.Query("SELECT * FROM movies")
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var movies []*Movie
	for row.Next() {
		var movie Movie
		err := row.Scan(&movie.ID, &movie.Title, &movie.ReleaseYear, &movie.Director)
		if err != nil {
			return nil, err
		}
		movies = append(movies, &movie)
	}
	return movies, nil
}

func (r *movieRepository) FindByID(id int) (*Movie, error) {
	// Query the database and return a movie with the specified ID
	row := r.db.QueryRow("SELECT * FROM movies WHERE id = ?", id)
	var movie Movie
	err := row.Scan(&movie.ID, &movie.Title, &movie.ReleaseYear, &movie.Director)
	if err != nil {
		return nil, err
	}
	return &movie, nil
}

func (r *movieRepository) Create(movie *Movie) error {
	// Insert a new movie into the database
	stmt := "INSERT INTO movies (title, release_year, director) VALUES (?, ?, ?)"
	_, err := r.db.Exec(stmt, movie.Title, movie.ReleaseYear, movie.Director)
	if err != nil {
		return err
	}
	return nil
}

func (r *movieRepository) Update(movie *Movie) error {
	// Update an existing movie in the database
	stmt := "UPDATE movies SET title = ?, release_year = ?, director = ? WHERE id = ?"
	_, err := r.db.Exec(stmt, movie.Title, movie.ReleaseYear, movie.Director, movie.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *movieRepository) Delete(id int) error {
	// Delete a movie from the database
	stmt := "DELETE FROM movies WHERE id = ?"
	_, err := r.db.Exec(stmt, id)
	if err != nil {
		return err
	}
	return nil
}

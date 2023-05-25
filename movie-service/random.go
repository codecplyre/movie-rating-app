package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/gofrs/uuid"
)

// create a random user name using uuid
func RandomDirector() string {
	return "director" + strconv.Itoa(rand.Intn(100))
}

// create a random movie name
func RandomMovie() string {
	return "movie" + uuid.Must(uuid.NewV4()).String()
}

// create a random movie release date and return time.Time format
func RandomReleaseDate() time.Time {
	year := rand.Intn(23) + 2000
	month := rand.Intn(12) + 1
	day := rand.Intn(28) + 1
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

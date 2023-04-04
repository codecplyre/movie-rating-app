-- Active: 1680618150350@@127.0.0.1@3306@movie_service
CREATE TABLE IF NOT EXISTS movies (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL UNIQUE,
    release_date DATE NOT NULL,
    director VARCHAR(255) NOT NULL
);
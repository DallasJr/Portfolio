package config

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectDatabase() {
	var err error
	DB, err = sql.Open("sqlite3", "data.db")
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	// Ping the DB
	if err = DB.Ping(); err != nil {
		panic(fmt.Sprintf("Failed to ping the database: %v", err))
	}

	_, err = createPortfolioTable()
	if err != nil {
		panic(fmt.Sprintf("Failed to create portfolio table: %v", err))
	}

	_, err = createMadeMoviesTable()
	if err != nil {
		panic(fmt.Sprintf("Failed to create made movies table: %v", err))
	}

	_, err = createPlayedMoviesTable()
	if err != nil {
		panic(fmt.Sprintf("Failed to create played movies table: %v", err))
	}

	fmt.Println("Database connection successful.")
}

func createPlayedMoviesTable() (sql.Result, error) {
	sqlStmt := `
    CREATE TABLE IF NOT EXISTS playedMovies (
        title TEXT PRIMARY KEY,
        releaseDate TEXT NOT NULL,
        description TEXT,
        image TEXT
    );
    `
	return DB.Exec(sqlStmt)
}

func createMadeMoviesTable() (sql.Result, error) {
	sqlStmt := `
    CREATE TABLE IF NOT EXISTS madeMovies (
        title TEXT PRIMARY KEY,
        releaseDate TEXT NOT NULL,
        description TEXT,
        image TEXT
    );
    `
	return DB.Exec(sqlStmt)
}

func createPortfolioTable() (sql.Result, error) {
	sqlStmt := `
    CREATE TABLE IF NOT EXISTS portfolios (
        name TEXT PRIMARY KEY,
        country TEXT NOT NULL,
        age INTEGER NOT NULL,
        description TEXT
    );
    `
	return DB.Exec(sqlStmt)
}

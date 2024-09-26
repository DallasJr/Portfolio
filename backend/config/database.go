package config

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"portfolio/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("portfolio.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database!")
	}
	log.Println("Database connected successfully.")

	// Migrate the schema
	DB.AutoMigrate(&models.Portfolio{}, &models.MadeMovie{}, &models.PlayedMovie{})
	// Ping the DB
	/*if err = DB.Ping(); err != nil {
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

	err = insertPortfolio()
	if err != nil {
		panic(fmt.Sprintf("Failed to insert portfolio: %v", err))
	}*/

	fmt.Println("Database connection successful.")
}

/*func createPlayedMoviesTable() (sql.Result, error) {
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

func insertPortfolio() error {
	sqlStmt := `INSERT INTO portfolios (name, country, age, description)
	VALUES ('je ne sais', 'france', 19, 'acteur');
	`
	_, err := DB.Exec(sqlStmt)
	return err
}
*/

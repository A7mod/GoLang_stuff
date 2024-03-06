package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB // lowercase means it is private. can't be accessed from package outside db
}

func NewDatabase() (*Database, error) {
	db, err := sql.Open("postgres", "postgresql://root:password@localhost:5433/go-chat?sslmode=disable")
	if err != nil {
		return nil, err
	}
	return &Database{db: db}, nil // return database struct that contains db, instantiated above, and nil in case of error
}

func (d *Database) Close() { // method to close the Database; as our DB is encapsulated.
	d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}

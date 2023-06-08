package persistence

import (
	"database/sql"
	"log"
	"os"

	"github.com/julianolorenzato/choosely/domain/poll"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func EstablishPostgresConnection() (*sql.DB, error) {
	// Open database's poll of connections
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL")+"?sslmode=disable")
	if err != nil {
		return nil, err
	}

	// Test database's poll of connections
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Get the database driver
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	// Create a new migrator
	m, err := migrate.NewWithDatabaseInstance("file://migrations", "main", driver)
	if err != nil {
		return nil, err
	}

	// Perform "up" migrations
	m.Up()

	log.Println("database initialised and migrations performed")

	return db, err
}

type PostgresPollRepository struct {
	writer *sql.DB
	reader *sql.DB
}

func NewPostgresPostgresPollRepository(w, r *sql.DB) *PostgresPollRepository {
	return &PostgresPollRepository{
		writer: w,
		reader: r,
	}
}

func (repo *PostgresPollRepository) GetByID(ID string) (*poll.Poll, error) {
	rows, err := repo.reader.Query("SELECT * FROM polls WHERE id = ?", ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var poll *poll.Poll

	for rows.Next() {
		err := rows.Scan(&poll.ID, &poll.Question, &poll.NumberOfChoices)
		if err != nil {
			return nil, err
		}
	}

	return poll, nil
}

func (repo *PostgresPollRepository) Save(poll *poll.Poll) error {
	return nil
}

func (repo *PostgresPollRepository) Create(poll *poll.Poll) error {
	_, err := repo.writer.Exec(
		`INSERT INTO polls
		(id, question, number_of_choices, options, votes, is_permanent, expires_at, created_at)
		VALUES
		(?, ?, ?, ?, ?, ?, ?, ?)`,
		poll.ID, poll.Question, poll.NumberOfChoices, poll.Options, poll.Votes, poll.IsPermanent, poll.ExpiresAt, poll.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

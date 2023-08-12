package persistence

import (
	"database/sql"
	"github.com/julianolorenzato/choosely/core/domain"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func establishPostgresConnection() (*sql.DB, error) {
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

// --------------------------------------------------------------------------------

type PostgresPollDB struct {
	writer *sql.DB
	reader *sql.DB
}

func NewPostgresPollDB() *PostgresPollDB {
	conn, err := establishPostgresConnection()
	if err != nil {
		log.Fatal(err)
	}

	return &PostgresPollDB{
		writer: conn,
		reader: conn,
	}
}

func (db *PostgresPollDB) GetByID(ID string) (*domain.Poll, error) {
	rows, err := db.reader.Query("SELECT * FROM polls WHERE id = $1", ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var poll *domain.Poll

	for rows.Next() {
		err := rows.Scan(&poll.ID, &poll.Question, &poll.NumberOfChoices)
		if err != nil {
			return nil, err
		}
	}

	return poll, nil
}

func (db *PostgresPollDB) Create(poll *domain.Poll) error {
	_, err := db.writer.Exec(
		`INSERT INTO polls
		(id, question, number_of_choices, options, is_permanent, expires_at, created_at)
		VALUES
		($1, $2, $3, $4, $5, $6, $7, $8)`,
		poll.ID, poll.Question, poll.NumberOfChoices, poll.Options, poll.IsPermanent, poll.ExpiresAt, poll.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (db *PostgresPollDB) Save(poll *domain.Poll) error {
	return nil
}

// --------------------------------------------------------------------------------

type PostgresVoteDB struct {
	writer *sql.DB
	reader *sql.DB
}

func NewPostgresVoteDB() *PostgresVoteDB {
	conn, err := establishPostgresConnection()
	if err != nil {
		log.Fatal(err)
	}

	return &PostgresVoteDB{
		writer: conn,
		reader: conn,
	}
}

func (db *PostgresVoteDB) GetByID(ID string) (*domain.Vote, error) {
	row := db.reader.QueryRow(`SELECT * FROM votes WHERE id = $1`, ID)

	var vote *domain.Vote

	err := row.Scan(&vote.ID, &vote.VoterID, &vote.PollID)
	if err != nil {
		return nil, err
	}

	return vote, nil
}

func (db *PostgresVoteDB) Create(vote *domain.Vote) error {
	_, err := db.writer.Exec(
		`INSERT INTO votes (id, voter_id, poll_id, choosen_options, created_at)
		VALUES ($1, $2, $3, $4, $5)`,
		vote.ID, vote.VoterID, vote.VoterID, vote.ChoosenOptions, vote.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (db *PostgresVoteDB) GetResults(pollID string) (map[string]uint, error) {
	rows, err := db.reader.Query(
		`SELECT option, COUNT(*) FROM votes,
		UNNEST(choosen_options) AS option
		WHERE poll_id = $1
		GROUP BY option`,
		pollID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := make(map[string]uint)

	for rows.Next() {
		var option string
		var votes uint

		err := rows.Scan(option, votes)
		if err != nil {
			return nil, err
		}

		results[option] = votes
	}

	return results, nil
}

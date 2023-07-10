package persistence

import (
	"database/sql"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/julianolorenzato/choosely/domain/poll"
	"github.com/julianolorenzato/choosely/domain/vote"
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

// --------------------------------------------------------------------------------

type PostgresPollRepository struct {
	writer *sql.DB
	reader *sql.DB
}

func NewPostgresPollRepository(w, r *sql.DB) *PostgresPollRepository {
	return &PostgresPollRepository{
		writer: w,
		reader: r,
	}
}

func (repo *PostgresPollRepository) GetByID(ID string) (*poll.Poll, error) {
	rows, err := repo.reader.Query("SELECT * FROM polls WHERE id = $1", ID)
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

func (repo *PostgresPollRepository) Create(poll *poll.Poll) error {
	_, err := repo.writer.Exec(
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

func (repo *PostgresPollRepository) Save(poll *poll.Poll) error {
	return nil
}

// --------------------------------------------------------------------------------

type PostgresVoteRepository struct {
	writer *sql.DB
	reader *sql.DB
}

func NewPostgresVoteRepository(w, r *sql.DB) *PostgresVoteRepository {
	return &PostgresVoteRepository{
		writer: w,
		reader: r,
	}
}

func (repo *PostgresVoteRepository) GetByID(ID string) (*vote.Vote, error) {
	row := repo.reader.QueryRow(`SELECT * FROM votes WHERE id = $1`, ID)

	var vote *vote.Vote

	err := row.Scan(&vote.ID, &vote.VoterID, &vote.PollID)
	if err != nil {
		return nil, err
	}

	return vote, nil
}

func (repo *PostgresVoteRepository) Create(vote *vote.Vote) error {
	_, err := repo.writer.Exec(
		`INSERT INTO votes (id, voter_id, poll_id, choosen_options, created_at)
		VALUES ($1, $2, $3, $4, $5)`,
		vote.ID, vote.VoterID, vote.VoterID, vote.ChoosenOptions, vote.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (repo *PostgresVoteRepository) GetResults(pollID string) (map[string]uint, error) {
	rows, err := repo.reader.Query(
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

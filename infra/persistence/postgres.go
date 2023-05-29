package persistence

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/julianolorenzato/choosely/domain/poll"
)

func InitialisePostgres() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_CONNECTION_STRING"))
	if err != nil {
		return nil, err
	}

	return db, err
}

type PollRepository struct {
	writer *sql.DB
	reader *sql.DB
}

func (repo *PollRepository) GetAllPolls() error {
	rows, err := repo.reader.Query("SELECT * FROM polls")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var poll poll.Poll

		err := rows.Scan(&poll.ID)
		if err != nil {
			return err
		}

		fmt.Println(rows)
	}

	return nil
}

func (repo *PollRepository) Create(poll poll.Poll) error {
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

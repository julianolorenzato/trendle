package persistence

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/julianolorenzato/choosely/domain/poll"
	_ "github.com/lib/pq"
)

func EstablishPostgresConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

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

func (repo *PostgresPollRepository) GetByID(ID string) *poll.Poll {
	rows /* err */, _ := repo.reader.Query("SELECT * FROM polls")
	// if err != nil {
	// 	return err
	// }
	defer rows.Close()

	for rows.Next() {
		var poll poll.Poll

		/* err :=  */
		rows.Scan(&poll.ID)
		// if err != nil {
		// 	return err
		// }

		fmt.Println(rows)
	}

	return nil
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

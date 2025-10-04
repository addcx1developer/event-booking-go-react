package repositories

import (
	"database/sql"

	"github.com/addcx1developer/event-booking-go-react/internal/models"
)

type EventStore struct {
	db *sql.DB
}

func New(db *sql.DB) *EventStore {
	return &EventStore{
		db: db,
	}
}

func (s *EventStore) Save(e *models.Event) error {
	query := `
		INSERT INTO events (name, description, location, date_time, user_id)
		VALUES (?, ?, ?, ?, ?)
	`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	e.ID = id

	return nil
}

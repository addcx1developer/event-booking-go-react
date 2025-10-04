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

func (s *EventStore) GetAll() ([]models.Event, error) {
	query := "SELECT * FROM events"

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []models.Event
	for rows.Next() {
		var event models.Event

		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func (s *EventStore) GetByID(id int64) (*models.Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := s.db.QueryRow(query, id)

	var event models.Event

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

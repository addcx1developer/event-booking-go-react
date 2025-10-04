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

func (s *EventStore) Save(e *models.Event) {}

package repositories

import (
	"database/sql"

	"github.com/addcx1developer/event-booking-go-react/internal/models"
	"github.com/addcx1developer/event-booking-go-react/internal/utils"
)

type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (s *UserStore) Save(u *models.User) error {
	query := `
		INSERT INTO users (email, password)
		VALUES (?, ?)
	`

	hashedPassword, err := utils.HashedPassword(u.Password)
	if err != nil {
		return err
	}

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = id

	return nil
}

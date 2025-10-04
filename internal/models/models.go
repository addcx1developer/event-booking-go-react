package models

import "time"

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	DateTime    time.Time `json:"date_time"`
	UserID      int64     `json:"user_id"`
}

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

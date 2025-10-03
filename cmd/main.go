package main

import (
	"fmt"
	"log"

	"github.com/addcx1developer/event-booking-go-react/internal/db"
)

func main() {
	cfg := config{
		maxOpenConns: 30,
		maxIdleConns: 30,
		maxIdleTime:  "15m",
	}

	// Database connection
	db, err := db.New(cfg.maxOpenConns, cfg.maxIdleConns, cfg.maxIdleTime)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	log.Println("Database connection pool established")

	app := &application{
		config: cfg,
	}

	mux := app.mount()
	if err := app.run(mux); err != nil {
		fmt.Println("err connecting")
	}

	log.Println(app.run(mux))
}

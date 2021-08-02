package storage

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

type Storage struct {
	config *Config
	sync.RWMutex
	DB *sqlx.DB
}

func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

func (s *Storage) Open() error {
	db, err := sqlx.Connect("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.DB = db

	return nil
}

func (s *Storage) Close() {
	err := s.DB.Close()

	if err != nil {
		log.Fatal(err)
	}
}

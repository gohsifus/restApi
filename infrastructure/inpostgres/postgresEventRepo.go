package inpostgres

import (
	"database/sql"
	"restApi/domain/entity"
	"time"
)

// PostgresEventRepo реализация репозитория для хранения в postgres
type PostgresEventRepo struct {
	store *sql.DB
}

// NewPostgresEventRepo ...
func NewPostgresEventRepo(connectionString string) (*PostgresEventRepo, error) {
	conn, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	return &PostgresEventRepo{store: conn}, nil
}

// Create ...
func (p *PostgresEventRepo) Create(event *entity.Event) (*entity.Event, error) { return nil, nil }

// Update ...
func (p *PostgresEventRepo) Update(id int, event *entity.Event) error { return nil }

// Delete ...
func (p *PostgresEventRepo) Delete(id int) error { return nil }

// GetEventsByDateInterval ...
func (p *PostgresEventRepo) GetEventsByDateInterval(from, to time.Time) ([]entity.Event, error) {
	return nil, nil
}

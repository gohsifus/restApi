package repository

import (
	"restApi/domain/entity"
	"time"
)

// EventRepo интерфейс репозитория для работы с событиями
type EventRepo interface {
	Create(event *entity.Event) (*entity.Event, error)
	Update(id int, event *entity.Event) error
	Delete(id int) error
	GetEventsByDateInterval(from, to time.Time) ([]entity.Event, error)
}

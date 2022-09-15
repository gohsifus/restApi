// Package service слой сервиса
package service

import (
	"restApi/domain/entity"
	"restApi/domain/repository"
)

//Event интерфейс сервиса для работы с событиями
type Event interface{
	CreateEvent(date, name, description string) (*entity.Event, error)
	DeleteEvent(id int) error
	UpdateEvent(id, date, name, description string) error
	GetEventsForDay(from, to string) ([]entity.Event, error)
	GetEventsForWeek(from, to string) ([]entity.Event, error)
	GetEventsForMonth(dayFromTargetMonth string) ([]entity.Event, error)
}

//Service содержит все сервисы для работы приложения
type Service struct{
	Event
}

// NewService ...
func NewService(eventRepo repository.EventRepo) *Service{
	return &Service{
		NewCalendar(eventRepo),
	}
}
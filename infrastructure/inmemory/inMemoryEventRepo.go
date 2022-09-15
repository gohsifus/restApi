package inmemory

import (
	"errors"
	"restApi/domain/entity"
	"restApi/errs"
	"sync"
	"time"
)

// MemoryEventRepo реализация репозитория для хранения в памяти
type MemoryEventRepo struct {
	store     map[int]entity.Event
	currentID int
	sync.RWMutex
}

// NewInMemoryRepo вернет реализацию репозитория для хранения в памяти
func NewInMemoryRepo() (*MemoryEventRepo, error) {
	return &MemoryEventRepo{
		store:     make(map[int]entity.Event),
		currentID: 0,
	}, nil
}

// Create ...
func (im *MemoryEventRepo) Create(event *entity.Event) (*entity.Event, error) {
	if !event.Validate() {
		return nil, errs.New(entity.ErrValidationFailed, errs.IncorrectDataErr)
	}

	im.Lock()
	event.ID = im.currentID
	im.store[event.ID] = *event
	im.currentID++
	im.Unlock()

	return event, nil
}

// Update ...
func (im *MemoryEventRepo) Update(id int, event *entity.Event) error {
	im.Lock()
	old, ok := im.store[id]
	if ok {
		if !event.Date.IsZero() {
			old.Date = event.Date
		}

		if event.Name != "" {
			old.Name = event.Name
		}

		if event.Description != "" {
			old.Description = event.Description
		}
	} else {
		im.Unlock()
		return errs.New(errors.New("item not found"), errs.BusinessLogicErr)
	}
	im.store[id] = old
	im.Unlock()

	return nil
}

// Delete ...
func (im *MemoryEventRepo) Delete(id int) error {
	im.Lock()
	delete(im.store, id)
	im.Unlock()

	return nil
}

// GetEventsByDateInterval ...
func (im *MemoryEventRepo) GetEventsByDateInterval(from, to time.Time) ([]entity.Event, error) {
	events := []entity.Event{}
	if from.After(to) {
		return nil, errors.New("from must be before to")
	}

	im.RLock()
	for _, v := range im.store {
		if (v.Date.After(from) || v.Date.Equal(from)) && (v.Date.Before(to) || v.Date.Equal(to)) {
			events = append(events, v)
		}
	}
	im.RUnlock()

	return events, nil
}

package entity

import (
	"encoding/json"
	"errors"
	"time"
)

// ErrValidationFailed ...
var ErrValidationFailed = errors.New("validation failed")

// Event ...
type Event struct {
	ID          int       `json:"id"`
	Date        time.Time `json:"date"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

// NewEvent создаст сущность - событие
func NewEvent(date time.Time, name, description string) *Event {
	return &Event{
		Date:        date,
		Name:        name,
		Description: description,
	}
}

// Validate провалидирует объект события
func (e Event) Validate() bool {
	if e.Date.IsZero() || e.Name == "" || e.Description == "" || e.ID < 0 {
		return false
	}
	return true
}

// ToJSON ...
func (e Event) ToJSON() ([]byte, error) {
	return json.Marshal(e)
}

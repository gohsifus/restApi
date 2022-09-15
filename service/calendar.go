package service

import (
	"errors"
	"restApi/domain/entity"
	"restApi/domain/repository"
	"restApi/errs"
	"strconv"
	"time"
)

const layout = "2006-01-02"

//Calendar Реализация сервиса для управления событиями
type Calendar struct {
	repository repository.EventRepo
}

//NewCalendar вернет календарь - реализацию сервиса для работы с событиями
func NewCalendar(repo repository.EventRepo) Calendar {
	return Calendar{
		repository: repo,
	}
}

// CreateEvent ...
func (c Calendar) CreateEvent(date, name, description string) (*entity.Event, error) {
	/*
		Бизнес-логика не должна зависить от сервера, поэтому
		сюда передаются необработанные параметры в виде строк.
		Handlers не должны заниматься преобразованием данных.
	*/
	d, err := time.Parse(layout, date)
	if err != nil {
		return nil, errs.New(err, errs.IncorrectDataErr)
	}

	event := entity.NewEvent(d, name, description)

	if !event.Validate() {
		return nil, errs.New(errors.New("validation failed"), errs.BusinessLogicErr)
	}

	return c.repository.Create(event)
}

// DeleteEvent ...
func (c Calendar) DeleteEvent(id int) error {
	return c.repository.Delete(id)
}

// UpdateEvent ...
func (c Calendar) UpdateEvent(id, date, name, description string) error {
	ident, err := strconv.Atoi(id)
	if err != nil {
		return errs.New(err, errs.IncorrectDataErr)
	}

	var d time.Time
	if date != "" {
		d, err = time.Parse(layout, date)
		if err != nil {
			return errs.New(err, errs.IncorrectDataErr)
		}
	}

	event := entity.NewEvent(d, name, description)

	err = c.repository.Update(ident, event)
	if err != nil {
		return errs.Wrap(err)
	}

	return nil
}

// GetEventsForDay ...
func (c Calendar) GetEventsForDay(from, to string) ([]entity.Event, error) {
	if (from == "" || to == "") || from != to {
		return nil, errs.New(errors.New("\"from\" not equal \"to\" or empty"), errs.BusinessLogicErr)
	}

	day, err := time.Parse(layout, to)
	if err != nil {
		return nil, errs.New(err, errs.IncorrectDataErr)
	}

	return c.repository.GetEventsByDateInterval(day, day)
}

//GetEventsForWeek ...
func (c Calendar) GetEventsForWeek(from, to string) ([]entity.Event, error) {
	if from == "" || to == "" {
		return nil, errs.New(errors.New("\"from\" or \"to\" is empty"), errs.IncorrectDataErr)
	}

	fromDate, err := time.Parse(layout, from)
	if err != nil {
		return nil, errs.New(err, errs.IncorrectDataErr)
	}

	toDate, err := time.Parse(layout, to)
	if err != nil {
		return nil, errs.New(err, errs.IncorrectDataErr)
	}

	if toDate.Sub(fromDate) != time.Hour*144 {
		return nil, errs.New(errors.New("must be 7 days between \"from\" and \"to\" (xxxx-xx-05 - xxxx-xx-11)"), errs.IncorrectDataErr)
	}

	fromYear, fromWeek := fromDate.ISOWeek()
	toYear, toWeek := toDate.ISOWeek()

	if fromYear != toYear || fromWeek != toWeek {
		return nil, errs.New(errors.New("\"from\" and \"to\" must belong the same week"), errs.IncorrectDataErr)
	}

	return c.repository.GetEventsByDateInterval(fromDate, toDate)
}

// GetEventsForMonth ...
func (c Calendar) GetEventsForMonth(dayFromTargetMonth string) ([]entity.Event, error) {
	if dayFromTargetMonth == "" {
		return nil, errs.New(errors.New("\"date\" must be not empty"), errs.IncorrectDataErr)
	}

	date, err := time.Parse(layout, dayFromTargetMonth)
	if err != nil {
		return nil, errs.New(err, errs.IncorrectDataErr)
	}

	fromDate := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.UTC)
	toDate := time.Date(date.Year(), date.Month()+1, 0, 0, 0, 0, 0, time.UTC)

	return c.repository.GetEventsByDateInterval(fromDate, toDate)
}

package handler

import (
	"encoding/json"
	"net/http"
	"net/url"
	"restApi/errs"
	"restApi/service"
	"strconv"
)

// Handler содержит все обработчики
type Handler struct {
	service *service.Service
}

// NewHandler ...
func NewHandler(service *service.Service) Handler {
	return Handler{
		service: service,
	}
}

// response функция для формирования ответа на запрос
func response(w http.ResponseWriter, status int, data interface{}, success bool) {
	w.WriteHeader(status)

	respData := make(map[string]interface{})

	if success {
		respData["result"] = data
	} else {
		respData["error"] = data
	}

	bytes, err := json.Marshal(respData)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("error: " + err.Error()))
		return
	}

	w.Write(bytes)
}

// Hello ...
func (h Handler) Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!"))
}

// CreateEvent ...
func (h Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		params, err := parseParams(r)
		if err != nil {
			cErr := errs.Wrap(err)
			response(w, cErr.Status(), cErr.Error(), false)
			//s.log.Error(cErr.Error())
		}

		event, err := h.service.Event.CreateEvent(
			params.Get("date"),
			params.Get("name"),
			params.Get("description"),
		)

		if err != nil {
			cErr := errs.Wrap(err)
			response(w, cErr.Status(), cErr.Error(), false)
			//s.log.Error(cErr.Error())
		} else {
			response(w, 200, event, true)
		}
	}
}

// UpdateEvent ...
func (h Handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		params, err := parseParams(r)
		if err != nil {
			cErr := errs.Wrap(err)
			response(w, cErr.Status(), cErr.Error(), false)
			//s.log.Error(cErr.Error())
		}

		err = h.service.Event.UpdateEvent(
			params.Get("id"),
			params.Get("date"),
			params.Get("name"),
			params.Get("description"),
		)

		if err != nil {
			cErr := errs.Wrap(err)
			response(w, cErr.Status(), cErr.Error(), false)
			//s.log.Error(cErr.Error())
		} else {
			response(w, 200, "update success", true)
		}
	}
}

// DeleteEvent ...
func (h Handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		id, err := strconv.Atoi(r.PostForm.Get("id"))

		if err != nil {
			cErr := errs.New(err, errs.IncorrectDataErr)
			response(w, cErr.Status(), cErr.Error(), false)
			//s.log.Error(cErr.Error())
		} else {
			err = h.service.Event.DeleteEvent(id)
			if err != nil {
				cErr := errs.Wrap(err)
				response(w, cErr.Status(), cErr.Error(), false)
				//s.log.Error(cErr.Error())
			} else {
				response(w, 200, "delete success", true)
			}
		}
	}
}

// GetEventsForDay ...
func (h Handler) GetEventsForDay(w http.ResponseWriter, r *http.Request) {
	args := r.URL.Query()

	events, err := h.service.Event.GetEventsForDay(args.Get("from"), args.Get("to"))
	if err != nil {
		cErr := errs.Wrap(err)
		response(w, cErr.Status(), cErr.Error(), false)
		//s.log.Error(cErr.Error())
	} else {
		response(w, 200, events, true)
	}
}

// GetEventsForWeek ...
func (h Handler) GetEventsForWeek(w http.ResponseWriter, r *http.Request) {
	args := r.URL.Query()

	events, err := h.service.Event.GetEventsForWeek(args.Get("from"), args.Get("to"))
	if err != nil {
		cErr := errs.Wrap(err)
		response(w, cErr.Status(), cErr.Error(), false)
		//s.log.Error(cErr.Error())
	} else {
		response(w, 200, events, true)
	}
}

// GetEventsForMonth ...
func (h Handler) GetEventsForMonth(w http.ResponseWriter, r *http.Request) {
	args := r.URL.Query()

	events, err := h.service.Event.GetEventsForMonth(args.Get("date"))
	if err != nil {
		cErr := errs.Wrap(err)
		response(w, cErr.Status(), cErr.Error(), false)
		//s.log.Error(cErr.Error())
	} else {
		response(w, 200, events, true)
	}
}

// Для парсинга параметров метода на update и create
func parseParams(r *http.Request) (url.Values, error) {
	err := r.ParseForm()
	if err != nil {
		return nil, errs.New(err, errs.IncorrectDataErr)
	}

	return r.PostForm, nil
}

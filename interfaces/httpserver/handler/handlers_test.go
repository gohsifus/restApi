package handler

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"restApi/infrastructure/inmemory"
	"restApi/service"
	"strings"
	"testing"
)

func TestHandler_Hello(t *testing.T) {
	repo, _ := inmemory.NewInMemoryRepo()
	handler := http.HandlerFunc(NewHandler(service.NewService(repo)).Hello)

	mockW := httptest.NewRecorder()
	mockR, _ := http.NewRequest("GET", "/", nil)

	handler.ServeHTTP(mockW, mockR)

	assert.Equal(t, http.StatusOK, mockW.Code)
	assert.Equal(t, []byte("Hello!"), mockW.Body.Bytes())
}

func TestHandler_CreateEvent(t *testing.T) {
	repo, _ := inmemory.NewInMemoryRepo()
	handler := http.HandlerFunc(NewHandler(service.NewService(repo)).CreateEvent)

	testTable := []struct {
		name       string
		data       url.Values
		statusWant int
	}{
		{
			data: url.Values{
				"date":        []string{"2022-09-09"},
				"name":        []string{"qwe"},
				"description": []string{"asd'"},
			},
			statusWant: http.StatusOK,
		},
		{
			data: url.Values{
				"date":        []string{""},
				"name":        []string{"qwe"},
				"description": []string{"asd'"},
			},
			statusWant: http.StatusBadRequest,
		},
		{
			data: url.Values{
				"date":        []string{"2022-09-09"},
				"name":        []string{""},
				"description": []string{"asd'"},
			},
			statusWant: http.StatusServiceUnavailable,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockW := httptest.NewRecorder()
			mockR, _ := http.NewRequest("POST", "/create_event", strings.NewReader(testCase.data.Encode()))
			mockR.Header.Add("Content-Type", "application/x-www-form-urlencoded")

			handler.ServeHTTP(mockW, mockR)

			assert.Equal(t, testCase.statusWant, mockW.Code)
		})
	}
}

func TestHandler_GetEventsForDay(t *testing.T) {
	repo, _ := inmemory.NewInMemoryRepo()
	handler := http.HandlerFunc(NewHandler(service.NewService(repo)).GetEventsForDay)

	testTable := []struct {
		name       string
		data       url.Values
		statusWant int
	}{
		{
			name: "getEvents by 2022-09-09 OK",
			data: url.Values{
				"from": []string{"2022-09-09"},
				"to":   []string{"2022-09-09"},
			},
			statusWant: http.StatusOK,
		},
		{
			name: "getEvents with incorrect date",
			data: url.Values{
				"from": []string{"2022-09qwqwe"},
				"to":   []string{"2022-09qwqwe"},
			},
			statusWant: http.StatusBadRequest,
		},
		{
			name: "getEvents with logic incorrect date",
			data: url.Values{
				"from": []string{"2022-09-10"},
				"to":   []string{"2022-09-09"},
			},
			statusWant: http.StatusServiceUnavailable,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockW := httptest.NewRecorder()
			mockR, _ := http.NewRequest("GET", "/events_for_day?"+testCase.data.Encode(), nil)

			handler.ServeHTTP(mockW, mockR)

			assert.Equal(t, testCase.statusWant, mockW.Code)
		})
	}
}

func TestHandler_DeleteEvent(t *testing.T) {
	repo, _ := inmemory.NewInMemoryRepo()
	handler := http.HandlerFunc(NewHandler(service.NewService(repo)).DeleteEvent)

	testTable := []struct {
		name       string
		data       url.Values
		statusWant int
	}{
		{
			name: "delete by id OK",
			data: url.Values{
				"id": []string{"0"},
			},
			statusWant: http.StatusOK,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockW := httptest.NewRecorder()
			mockR, _ := http.NewRequest("POST", "/delete_event", strings.NewReader(testCase.data.Encode()))
			mockR.Header.Add("Content-Type", "application/x-www-form-urlencoded")

			handler.ServeHTTP(mockW, mockR)

			t.Log(mockW.Body.String())

			assert.Equal(t, testCase.statusWant, mockW.Code)
		})
	}
}

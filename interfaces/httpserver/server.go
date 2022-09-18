package httpserver

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"restApi/domain/repository"
	"restApi/interfaces/httpserver/configs"
	"restApi/interfaces/httpserver/handler"
	"restApi/logger"
	"restApi/metrics"
	"restApi/middleware"
	"restApi/service"
)

// Server http server
type Server struct {
	handler handler.Handler
	mux     *http.ServeMux
	config  *configs.ServerConfig
	log     *logger.Log
}

// NewServer ...
func NewServer(config *configs.ServerConfig, eventRepo repository.EventRepo) (*Server, error) {
	h := handler.NewHandler(service.NewService(eventRepo))
	mux := http.NewServeMux()
	logger, err := logger.NewLogger(config.PathToLog)
	if err != nil {
		return nil, err
	}

	return &Server{
		handler: h,
		mux:     mux,
		config:  config,
		log:     logger,
	}, nil
}

// Start запуск сервера
func (s Server) Start() {
	fmt.Println("start server")
	s.log.Info("start server")
	s.ConfigureServer()

	addr := s.config.Host + ":" + s.config.Port

	http.ListenAndServe(addr, s.mux)
}

// ConfigureServer сконфигурирует сервер, назначив обработчики
func (s Server) ConfigureServer() {
	m := metrics.NewMetrics()

	s.mux.Handle("/", middleware.Logging(http.HandlerFunc(s.handler.Hello), s.log, m))
	s.mux.Handle("/create_event", middleware.Logging(http.HandlerFunc(s.handler.CreateEvent), s.log, m))
	s.mux.Handle("/update_event", middleware.Logging(http.HandlerFunc(s.handler.UpdateEvent), s.log, m))
	s.mux.Handle("/delete_event", middleware.Logging(http.HandlerFunc(s.handler.DeleteEvent), s.log, m))
	s.mux.Handle("/events_for_day", middleware.Logging(http.HandlerFunc(s.handler.GetEventsForDay), s.log, m))
	s.mux.Handle("/events_for_week", middleware.Logging(http.HandlerFunc(s.handler.GetEventsForWeek), s.log, m))
	s.mux.Handle("/events_for_month", middleware.Logging(http.HandlerFunc(s.handler.GetEventsForMonth), s.log, m))
	s.mux.Handle("/metrics", promhttp.Handler())
}

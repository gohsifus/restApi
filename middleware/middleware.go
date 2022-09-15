package middleware

import (
	"fmt"
	"net/http"
	"restApi/logger"
)

// Logging оборачивает запрос, позволяя добавить логирование
func Logging(f http.Handler, logger *logger.Log) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Запрос: " + r.URL.String())
		logger.Info("Запрос: " + r.URL.String())
		f.ServeHTTP(w, r)
	})
}

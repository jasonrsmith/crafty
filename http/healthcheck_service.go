package http

import "net/http"

// HealthCheckHandler is a simple service check
type HealthCheckHandler struct{}

// NewHealthCheckHandler handles the health checks
func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}

func (*HealthCheckHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
}

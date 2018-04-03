package http

import (
	"encoding/json"
	"github.com/jasonrsmith/crafty"
	"log"
	"net/http"
	"strings"
)

const ErrInvalidJSON = crafty.Error("invalid json")

// Handler is a collection of all the service handlers.
type Handler struct {
	CraftyHandler      http.Handler
	HealthCheckHandler http.Handler
}

// ServeHTTP delegates a request to the appropriate subhandler.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/health") {
		h.HealthCheckHandler.ServeHTTP(w, r)
	} else {
		h.CraftyHandler.ServeHTTP(w, r)
	}
	log.Println(r.URL.Path)
}

// Error writes an API error message to the response and logger.
func Error(w http.ResponseWriter, err error, code int, logger *log.Logger) {
	logger.Printf("http error: %s (code=%d)", err, code)
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(&errorResponse{Err: err.Error()})
}

// errorResponse is a generic response for sending a error.
type errorResponse struct {
	Err string `json:"err,omitempty"`
}

// NotFound writes an API error message to the response.
func NotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{}` + "\n"))
}

// encodeJSON encodes v to w in JSON format. Error() is called if encoding fails.
func encodeJSON(w http.ResponseWriter, v interface{}, logger *log.Logger) {
	if err := json.NewEncoder(w).Encode(v); err != nil {
		Error(w, err, http.StatusInternalServerError, logger)
	}
}

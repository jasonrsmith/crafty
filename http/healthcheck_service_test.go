package http_test

import (
	"github.com/jasonrsmith/crafty/http"
	nethttp "net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	request, err := nethttp.NewRequest("GET", "/health", nil)
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	handler := http.NewHealthCheckHandler()
	handler.ServeHTTP(recorder, request)

	assert.Equal(t, nethttp.StatusOK, recorder.Code)
}

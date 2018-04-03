package http_test

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/jasonrsmith/crafty/http"
	"log"
	nethttp "net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockHandler struct {
	serveHTTP func(w nethttp.ResponseWriter, r *nethttp.Request)
}

func (mockHandler MockHandler) ServeHTTP(w nethttp.ResponseWriter, r *nethttp.Request) {
	mockHandler.serveHTTP(w, r)
}

func TestHealthURLGoesToHealthHandler(t *testing.T) {
	request, _ := nethttp.NewRequest("GET", "/health", nil)
	recorder := httptest.NewRecorder()

	var craftyHandlerCalled, healthCheckHandlerCalled bool

	craftyHandler := MockHandler{
		serveHTTP: func(w nethttp.ResponseWriter, r *nethttp.Request) {
			craftyHandlerCalled = true
		},
	}
	healthCheckHandler := MockHandler{
		serveHTTP: func(w nethttp.ResponseWriter, r *nethttp.Request) {
			healthCheckHandlerCalled = true
		},
	}

	handler := &http.Handler{
		CraftyHandler:      craftyHandler,
		HealthCheckHandler: healthCheckHandler,
	}

	handler.ServeHTTP(recorder, request)
	assert.True(t, healthCheckHandlerCalled)
	assert.False(t, craftyHandlerCalled)
}

func TestOtherURLsGoToCraftyHandler(t *testing.T) {
	request, _ := nethttp.NewRequest("GET", "/something", nil)
	recorder := httptest.NewRecorder()

	var craftyHandlerCalled, healthCheckHandlerCalled bool

	craftyHandler := MockHandler{
		serveHTTP: func(w nethttp.ResponseWriter, r *nethttp.Request) {
			craftyHandlerCalled = true
		},
	}
	healthCheckHandler := MockHandler{
		serveHTTP: func(w nethttp.ResponseWriter, r *nethttp.Request) {
			healthCheckHandlerCalled = true
		},
	}

	handler := &http.Handler{
		CraftyHandler:      craftyHandler,
		HealthCheckHandler: healthCheckHandler,
	}

	handler.ServeHTTP(recorder, request)
	assert.False(t, healthCheckHandlerCalled)
	assert.True(t, craftyHandlerCalled)
}

func TestErrorWritesCodeToHeader(t *testing.T) {
	var (
		expectedResponseCode = 123

		logBuf bytes.Buffer
		logger = log.New(&logBuf, "logger: ", log.Lshortfile)
	)

	recorder := httptest.NewRecorder()
	http.Error(recorder, errors.New("error"), expectedResponseCode, logger)

	assert.Equal(t, expectedResponseCode, recorder.Code)
}

func TestErrorResponseBody(t *testing.T) {
	var (
		responseMessage      = "theerror"
		expectedResponseBody = fmt.Sprintf(`{"err":"%s"}
`, responseMessage)

		logBuf bytes.Buffer
		logger = log.New(&logBuf, "logger: ", log.Lshortfile)
	)

	recorder := httptest.NewRecorder()
	http.Error(recorder, errors.New(responseMessage), 0, logger)
	assert.Equal(t, expectedResponseBody, recorder.Body.String())
}

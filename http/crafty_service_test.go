package http_test

import (
	"bytes"
	"github.com/jasonrsmith/crafty"
	"github.com/jasonrsmith/crafty/http"
	nethttp "net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockCraftyService struct {
	addPrimitive      func(prim crafty.Primitive) error
	findAllPrimitives func() ([]crafty.Primitive, error)
}

func (svc *MockCraftyService) AddPrimitive(prim crafty.Primitive) error {
	return svc.addPrimitive(prim)
}
func (svc *MockCraftyService) FindAllPrimitives() ([]crafty.Primitive, error) {
	return svc.findAllPrimitives()
}

func TestPostNoPrimitiveReturnsError(t *testing.T) {
	json := []byte(`{}`) //TODO: add empty json test/handling
	//json := []byte(`{"primitive": {"name": "A thing"}}`)
	request, _ := nethttp.NewRequest("POST", "/primitives", bytes.NewBuffer(json))
	recorder := httptest.NewRecorder()
	target := http.NewCraftyHandler(
		&MockCraftyService{},
	)

	target.ServeHTTP(recorder, request)
	assert.Equal(t, nethttp.StatusBadRequest, recorder.Code)
}

func TestPostPrimitiveReturnsCreatedResponse(t *testing.T) {
	json := []byte(`{"primitive": {"name": "A thing"}}`)
	request, _ := nethttp.NewRequest("POST", "/primitives", bytes.NewBuffer(json))
	recorder := httptest.NewRecorder()
	addPrimitiveCalled := false
	mock := &MockCraftyService{
		addPrimitive: func(prim crafty.Primitive) error {
			addPrimitiveCalled = true
			return nil
		},
	}

	target := http.NewCraftyHandler(mock)

	target.ServeHTTP(recorder, request)
	assert.Equal(t, nethttp.StatusCreated, recorder.Code)
	assert.True(t, addPrimitiveCalled)
}

func TestGetPrimitive(t *testing.T) {
	expectedJSON := `{"primitives":[{"id":"00000000-0000-0000-0000-000000000000","name":"A thing"}]}
`
	request, _ := nethttp.NewRequest("GET", "/primitives", nil)
	recorder := httptest.NewRecorder()
	target := http.NewCraftyHandler(
		&MockCraftyService{
			findAllPrimitives: func() ([]crafty.Primitive, error) {
				return []crafty.Primitive{
					{Name: "A thing"},
				}, nil
			},
		},
	)

	target.ServeHTTP(recorder, request)
	assert.Equal(t, nethttp.StatusOK, recorder.Code)
	assert.Equal(t, expectedJSON, recorder.Body.String())
}

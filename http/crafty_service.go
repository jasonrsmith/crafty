package http

import (
	"encoding/json"
	"jasmith/crafty"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	uuid "github.com/satori/go.uuid"
)

var unitializedUUID = uuid.UUID{}

// CraftyHandler represents an HTTP API handler for dials.
type CraftyHandler struct {
	*httprouter.Router

	CraftyService crafty.Service

	Logger *log.Logger
}

// NewCraftyHandler returns a new instance of CraftyHandler.
func NewCraftyHandler(service crafty.Service) *CraftyHandler {
	h := &CraftyHandler{
		Router:        httprouter.New(),
		CraftyService: service,
		Logger:        log.New(os.Stderr, "", log.LstdFlags),
	}
	h.POST("/primitives", h.handlePostPrimitives)
	h.GET("/primitives", h.handleGetPrimitives)
	return h
}

type getPrimitivesResponse struct {
	Primitives []crafty.Primitive `json:"primitives,omitempty"`
	Err        string             `json:"err,omitempty"`
}

// handleGetPrimitives handles requests to fetch a all primitives
func (h *CraftyHandler) handleGetPrimitives(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	prims, err := h.CraftyService.FindAllPrimitives()

	if err != nil {
		Error(w, err, http.StatusInternalServerError, h.Logger)
	} else if prims == nil {
		NotFound(w)
	} else {
		encodeJSON(w, &getPrimitivesResponse{Primitives: prims}, h.Logger)
	}
}

type postPrimitivesRequest struct {
	Primitive crafty.Primitive `json:"primitive"`
}

type postPrimitivesResponse struct {
	ID  uuid.UUID `json:"id,omitempty"`
	Err string    `json:"err,omitempty"`
}

// handlePostPrimitives handles requests to create new primitives
func (h *CraftyHandler) handlePostPrimitives(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Decode request.
	var req postPrimitivesRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		Error(w, ErrInvalidJSON, http.StatusBadRequest, h.Logger)
		return
	}
	prim := req.Primitive
	if prim.Name == "" {
		Error(w, ErrInvalidJSON, http.StatusBadRequest, h.Logger)
		return
	}
	if prim.ID == unitializedUUID {
		prim.ID = uuid.Must(uuid.NewV4())
	}

	// add primitive
	switch err := h.CraftyService.AddPrimitive(prim); err {
	case nil:
		w.WriteHeader(http.StatusCreated)
		encodeJSON(w, postPrimitivesResponse{ID: prim.ID}, h.Logger)
	default:
		Error(w, err, http.StatusInternalServerError, h.Logger)
	}
}

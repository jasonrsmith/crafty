package crafty

import uuid "github.com/satori/go.uuid"

// Primitive is an ingredient for a Composite, and can't be broken down
type Primitive struct {
	// TODO: test removing this breaks validation
	ID   uuid.UUID `json:"id,omitempty"`
	Name string    `json:"name"`
}

// CompositePart is a count of either a Primitive or a Composite
type CompositePart struct {
	Count int         `json:"count"`
	Part  interface{} `json:"part"`
}

// Composite is a grouping of CompositeParts, and represents a fully crafted item
type Composite struct {
	ID       uuid.UUID       `json:"id"`
	Name     string          `json:"name"`
	PartList []CompositePart `json:"parts"`
}

// Service manages crafts
type Service interface {
	AddPrimitive(Primitive) error
	FindAllPrimitives() ([]Primitive, error)
}

// Client creates a connection to service
// type Client interface {
// 	CraftyService() Service
// }

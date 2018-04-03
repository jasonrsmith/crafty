package memstore_test

import (
	"github.com/jasonrsmith/crafty"
	"github.com/jasonrsmith/crafty/memstore"
	"testing"

	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestAddOne(t *testing.T) {
	svc := memstore.NewCraftyService()
	primName := "prim"
	prim := crafty.Primitive{
		ID:   uuid.Must(uuid.NewV4()),
		Name: primName,
	}
	svc.AddPrimitive(prim)
	primList, err := svc.FindAllPrimitives()
	assert.Equal(t, primName, primList[0].Name)
	assert.Nil(t, err)
}

func TestAddTwo(t *testing.T) {
	svc := memstore.NewCraftyService()
	primName1 := "prim1"
	primName2 := "prim2"
	prim1 := crafty.Primitive{
		ID:   uuid.Must(uuid.NewV4()),
		Name: primName1,
	}
	prim2 := crafty.Primitive{
		ID:   uuid.Must(uuid.NewV4()),
		Name: primName2,
	}
	svc.AddPrimitive(prim1)
	svc.AddPrimitive(prim2)
	primitiveList, err := svc.FindAllPrimitives()
	assert.Equal(t, primName1, primitiveList[0].Name)
	assert.Equal(t, primName2, primitiveList[1].Name)
	assert.Nil(t, err)
}

package memstore_test

import (
	"jasmith/crafthelper/memstore"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddOne(t *testing.T) {
	svc := memstore.NewCraftHelperService()
	primName := "prim"
	svc.AddPrimitive(primName)
	primitiveList, err := svc.FindAllPrimitives()
	assert.Equal(t, primName, primitiveList[0].Name)
	assert.Nil(t, err)
}

func TestAddTwo(t *testing.T) {
	svc := memstore.NewCraftHelperService()
	primName1 := "prim1"
	primName2 := "prim2"
	svc.AddPrimitive(primName1)
	svc.AddPrimitive(primName2)
	primitiveList, err := svc.FindAllPrimitives()
	assert.Equal(t, primName1, primitiveList[0].Name)
	assert.Equal(t, primName2, primitiveList[1].Name)
	assert.Nil(t, err)
}

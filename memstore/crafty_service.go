package memstore

import (
	"github.com/jasonrsmith/crafty"
)

var _ crafty.Service = &CraftyService{}

type CraftyService struct {
	primitives map[string]crafty.Primitive
}

func NewCraftyService() *CraftyService {
	return &CraftyService{
		make(map[string]crafty.Primitive),
	}
}

func (svc *CraftyService) AddPrimitive(prim crafty.Primitive) error {
	svc.primitives[prim.Name] = prim
	return nil
}

func (svc *CraftyService) FindAllPrimitives() ([]crafty.Primitive, error) {
	primList := []crafty.Primitive{}
	for _, prim := range svc.primitives {
		primList = append(primList, prim)
	}
	return primList, nil
}

package memstore

import "jasmith/crafthelper"

type CraftHelperService struct {
	primitives map[string]crafthelper.Primitive
}

func NewCraftHelperService() *CraftHelperService {
	return &CraftHelperService{
		make(map[string]crafthelper.Primitive),
	}
}

func (svc *CraftHelperService) AddPrimitive(name string) error {
	svc.primitives[name] = crafthelper.Primitive{name}
	return nil
}

func (svc *CraftHelperService) FindAllPrimitives() ([]crafthelper.Primitive, error) {
	primList := []crafthelper.Primitive{}
	for _, prim := range svc.primitives {
		primList = append(primList, prim)
	}
	return primList, nil
}

package crafthelper

type Primitive struct {
	Name string `json:"name"`
}

/*
type CompositeElement struct {
	Count int
	Element interface{}
}

type Composite struct {
	ElementList []CompositeElement
}
*/

type Client interface {
	CraftHelperService() CraftHelperService
}

type CraftHelperService interface {
	AddPrimitive(name string) error
	FindAllPrimitives() ([]Primitive, error)
}

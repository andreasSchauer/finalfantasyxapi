package servicetypes

type ServiceParams interface {
	AlBhedParams | TurnOrderParams | DropChanceParams | CritChanceParams
}

type ServiceResponse interface {
	AlBhedResponse | TurnOrderResponse | DropChanceResponse | CritChanceResponse
}

type SampleParams[P ServiceParams] struct {
	Description string `json:"description"`
	Params      P      `json:"params"`
}

type NonDuplicable interface {
	GetDuplicateKey() string
}

func safeDuplicateKey[T NonDuplicable](objPtr *T) string {
	if objPtr == nil {
		return "null"
	}

	obj := *objPtr
	return obj.GetDuplicateKey()
}

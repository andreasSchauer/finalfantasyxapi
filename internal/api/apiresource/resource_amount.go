package apiresource

type NameAmount interface {
	GetName() string
	GetVersion() *int32
	GetVal() int32
}

// ResourceAmountType can be converted into a generic ResourceAmount, but also has another, more descriptive alias (e.g. BaseStat)
type ResourceAmountType[A APIResource] interface {
	ToResAmount() ResourceAmount[A]
	NameAmount
}

type ResourceAmount[A APIResource] struct {
	Resource A     `json:"resource"`
	Amount   int32 `json:"amount"`
}

func (ra ResourceAmount[A]) GetAPIResource() APIResource {
	return ra.Resource
}

func NewResourceAmount[A APIResource](resource A, amount int32) ResourceAmount[A] {
	return ResourceAmount[A]{
		Resource: resource,
		Amount:   amount,
	}
}
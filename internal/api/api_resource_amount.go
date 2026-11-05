package api


type ResourceAmount[A APIResource] struct {
	Resource A     `json:"resource"`
	Amount   int32 `json:"amount"`
}

func (ra ResourceAmount[A]) GetAPIResource() APIResource {
	return ra.Resource
}

func newResourceAmount[A APIResource](resource A, amount int32) ResourceAmount[A] {
	return ResourceAmount[A]{
		Resource: resource,
		Amount:   amount,
	}
}

func GetResAmtTypeMap[T ResourceAmountType[A], A APIResource](items []T) map[string]int32 {
	resAmts := ResAmtTypesToStructs(items)
	return GetResAmtMap(resAmts)
}

func GetResAmtMap[A APIResource](items []ResourceAmount[A]) map[string]int32 {
	amountMap := make(map[string]int32)

	for _, item := range items {
		key := item.Resource.GetTestKey()
		amountMap[key] = item.Amount
	}

	return amountMap
}

func GetResAmtIDMap[A APIResource](items []ResourceAmount[A]) map[int32]int32 {
	amountMap := make(map[int32]int32)

	for _, item := range items {
		url := item.Resource.GetURL()
		id := getIDFromURL(url)
		amountMap[id] = item.Amount
	}

	return amountMap
}

func ResAmtTypesToStructs[T ResourceAmountType[A], A APIResource](items []T) []ResourceAmount[A] {
	resAmts := []ResourceAmount[A]{}

	for _, item := range items {
		resAmts = append(resAmts, item.ToResAmount())
	}

	return resAmts
}

func resAmtMapToFloat(intMap map[string]int32) map[string]float64 {
	floatMap := make(map[string]float64)

	for key := range intMap {
		float := float64(intMap[key])
		floatMap[key] = float
	}

	return floatMap
}
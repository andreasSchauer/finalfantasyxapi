package handlerinput

import (
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
)

func GetResAmtTypeMap[T apiresource.ResourceAmountType[A], A apiresource.APIResource](items []T) map[string]int32 {
	resAmts := ResAmtTypesToStructs(items)
	return GetResAmtMap(resAmts)
}

func GetResAmtMap[A apiresource.APIResource](items []apiresource.ResourceAmount[A]) map[string]int32 {
	amountMap := make(map[string]int32, len(items))

	for _, item := range items {
		key := item.Resource.GetTestKey()
		amountMap[key] = item.Amount
	}

	return amountMap
}

func GetResAmtIDMap[A apiresource.APIResource](items []apiresource.ResourceAmount[A]) map[int32]int32 {
	amountMap := make(map[int32]int32, len(items))

	for _, item := range items {
		url := item.Resource.GetURL()
		id := ah.GetIDFromURL(url)
		amountMap[id] = item.Amount
	}

	return amountMap
}

func ResAmtTypesToStructs[T apiresource.ResourceAmountType[A], A apiresource.APIResource](items []T) []apiresource.ResourceAmount[A] {
	resAmts := make([]apiresource.ResourceAmount[A], len(items))

	for i, item := range items {
		resAmts[i] = item.ToResAmount()
	}

	return resAmts
}

func ResAmtMapToFloat(intMap map[string]int32) map[string]float64 {
	floatMap := make(map[string]float64, len(intMap))

	for key := range intMap {
		float := float64(intMap[key])
		floatMap[key] = float
	}

	return floatMap
}

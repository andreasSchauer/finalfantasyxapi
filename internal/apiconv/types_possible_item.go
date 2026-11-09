package apiconv

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	th "github.com/andreasSchauer/finalfantasyxapi/internal/typehelpers"
)

func ConvertPossibleItem(item seedtypes.PossibleItem) apitypes.PossibleItem {
	return th.NewPossibleItem(item.ItemAmount, item.Chance)
}

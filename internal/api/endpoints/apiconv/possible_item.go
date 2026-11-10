package apiconv

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	th "github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/typehelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func ConvertPossibleItem(item seedtypes.PossibleItem) apitypes.PossibleItem {
	return th.NewPossibleItem(item.ItemAmount, item.Chance)
}

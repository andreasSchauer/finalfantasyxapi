package apiconv

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func ConvertBlitzballItem(bi seedtypes.BlitzballItem) apitypes.PossibleItem {
	return ConvertPossibleItem(bi.PossibleItem)
}

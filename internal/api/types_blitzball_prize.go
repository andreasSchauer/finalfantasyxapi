package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func ConvertBlitzballItem(bi seedtypes.BlitzballItem) apitypes.PossibleItem {
	return ConvertPossibleItem(bi.PossibleItem)
}

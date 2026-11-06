package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func convertAgilitySubtier(st seedtypes.AgilitySubtier) apitypes.AgilitySubtier {
	return apitypes.AgilitySubtier{
		FromAgility: st.MinAgility,
		ToAgility:   st.MaxAgility,
		CharMinICV:  st.CharacterMinICV,
	}
}

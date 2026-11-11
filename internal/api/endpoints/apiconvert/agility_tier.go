package apiconvert

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func AgilitySubtier(st seedtypes.AgilitySubtier) apitypes.AgilitySubtier {
	return apitypes.AgilitySubtier{
		FromAgility: st.MinAgility,
		ToAgility:   st.MaxAgility,
		CharMinICV:  st.CharacterMinICV,
	}
}

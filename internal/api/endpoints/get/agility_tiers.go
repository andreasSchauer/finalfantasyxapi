package get

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apiconv"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func GetAgilityTier(r *http.Request, i handlerinput.Main[seedtypes.AgilityTier, apitypes.AgilityTier, apiresource.Unnamed], id int32) (apitypes.AgilityTier, error) {
	agilityTier, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.AgilityTier{}, err
	}

	response := apitypes.AgilityTier{
		ID:          agilityTier.ID,
		FromAgility: agilityTier.MinAgility,
		ToAgility:   agilityTier.MaxAgility,
		TickSpeed:   agilityTier.TickSpeed,
		MonMinICV:   agilityTier.MonsterMinICV,
		MonMaxICV:   agilityTier.MonsterMaxICV,
		CharMaxICV:  agilityTier.CharacterMaxICV,
		CharMinICVs: ah.ConvertObjSlice(agilityTier.CharacterMinICVs, apiconv.ConvertAgilitySubtier),
	}

	return response, nil
}

package get

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/assemble"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func GetRonsoRage(r *http.Request, i handlerinput.Main[seedtypes.RonsoRage, apitypes.RonsoRage, apiresource.Named], id int32) (apitypes.RonsoRage, error) {
	ronsoRage, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.RonsoRage{}, err
	}

	rel, err := assemble.GetRonsoRageRelationships(r, ronsoRage)
	if err != nil {
		return apitypes.RonsoRage{}, err
	}

	response := apitypes.RonsoRage{
		ID:          ronsoRage.ID,
		Name:        ronsoRage.Name,
		Description: ronsoRage.Description,
		Effect:      ronsoRage.Effect,
		Overdrive:   handlerinput.NameToNamedAPIResource(apilookups.E.Overdrives, ronsoRage.Name, nil),
		Monsters:    rel.Monsters,
	}

	return response, nil
}

func RetrieveRonsoRages(r *http.Request, i handlerinput.Main[seedtypes.RonsoRage, apitypes.RonsoRage, apiresource.Named]) ([]int32, error) {
	return i.VerifyQueriesRetrieve(r)
}

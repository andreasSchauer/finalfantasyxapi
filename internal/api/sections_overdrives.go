package api

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/sectiontypes"
)


func createOverdriveSimple(_ *http.Request, id int32, _ Subsection) (apiresource.SimpleResource, error) {
	i := E.Overdrives
	overdrive, _ := lookup.GetResourceByID(id, i.objLookupID)

	overdriveSimple := sectiontypes.OverdriveSimple{
		ID:                 overdrive.ID,
		URL:                ah.CreateResourceURL(i.Endpoint, id),
		Name:               overdrive.Name,
		Version:            overdrive.Version,
		Specification:      overdrive.Specification,
		Rank:               overdrive.Rank,
		OverdriveAbilities: ah.ConvertObjSlice(overdrive.OverdriveAbilities, abilityRefString),
	}

	return overdriveSimple, nil
}

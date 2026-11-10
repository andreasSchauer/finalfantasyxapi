package sections

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/sections/sectiontypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
)

func CreateOverdriveSimple(_ *http.Request, id int32, _ handlerinput.Subsection) (apiresource.SimpleResource, error) {
	i := apilookups.E.Overdrives
	overdrive, _ := lookup.GetResourceByID(id, i.ObjLookupID)

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

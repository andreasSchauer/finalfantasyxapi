package api

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/sectiontypes"
)

func createSubquestSimple(_ *http.Request, id int32, _ handlerinput.Subsection) (apiresource.SimpleResource, error) {
	i := apilookups.E.Subquests
	subquest, _ := lookup.GetResourceByID(id, i.ObjLookupID)

	subquestSimple := sectiontypes.SubquestSimple{
		ID:        subquest.ID,
		URL:       ah.CreateResourceURL(i.Endpoint, id),
		Name:      subquest.Name,
		Condition: subquest.Completion.Condition,
		Areas:     locAreaStrings(subquest.Completion.Areas),
		Reward:    convertItemAmountSimple(subquest.GetItemAmount()),
	}

	return subquestSimple, nil
}
